// WHAT: 下载器核心逻辑 — Goroutine 工作池模式
// WHY: 这是 Go 并发模型经典应用，综合运用 goroutine + channel + context + 接口
// CONTRAST:
//   - TS:   用 Promise.all + fetch，单线程事件循环（但非阻塞 I/O）
//   - Rust: 用 tokio::spawn + mpsc channel，语义类似但多了生命周期管理
//   - Go:   最简洁的并发写法，不需要 async/await 关键字

package downloader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Downloader 并发下载器
// WHAT: 包含所有下载逻辑的结构体
// WHY: 把状态和行为封装在一起（Go 没有 class，用 struct + 方法）
type Downloader struct {
	// opts 配置（私有）
	opts Options
	// client HTTP 客户端（私有）
	client *http.Client
}

// NewDownloader 创建下载器（构造函数习惯用法）
func NewDownloader(opts Options) *Downloader {
	return &Downloader{
		opts: opts,
		client: &http.Client{
			Timeout: opts.Timeout,
		},
	}
}

// Download 执行批量下载（导出方法，主要 API）
// ctx: 用于取消所有下载
// tasks: 下载任务列表
// reporter: 进度报告器（接受接口！遵循"接受接口"原则）
// 返回：每个任务的结果（保持顺序）
func (d *Downloader) Download(ctx context.Context, tasks []DownloadTask, reporter ProgressReporter) []DownloadResult {
	// 创建 channel
	// jobCh: 任务分发 channel（带缓冲，避免阻塞）
	jobCh := make(chan DownloadTask, len(tasks))
	resultCh := make(chan DownloadResult, len(tasks))

	// 启动 worker goroutine 池
	var wg sync.WaitGroup
	for i := 0; i < d.opts.Concurrency; i++ {
		wg.Add(1)
		go d.worker(ctx, i, jobCh, resultCh, reporter, &wg)
	}

	// 分发任务
	for _, task := range tasks {
		jobCh <- task
	}
	close(jobCh) // 关闭任务 channel，workers 会退出

	// 等待所有 worker 完成（在单独的 goroutine 中）
	go func() {
		wg.Wait()
		close(resultCh) // 所有 worker 完成后关闭结果 channel
	}()

	// 收集结果
	results := make([]DownloadResult, 0, len(tasks))
	for result := range resultCh {
		results = append(results, result)
	}
	return results
}

// worker 工作协程（私有方法）
// WHAT: 从 jobCh 取任务，下载，结果写入 resultCh
// WHY: 这是 Go 并发中最经典的 worker pool 模式
func (d *Downloader) worker(ctx context.Context, id int, jobs <-chan DownloadTask, results chan<- DownloadResult, reporter ProgressReporter, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range jobs {
		// 检查上下文是否已取消
		select {
		case <-ctx.Done():
			results <- DownloadResult{Task: task, Error: ctx.Err()}
			return
		default:
		}

		// 执行下载（带重试）
		result := d.downloadWithRetry(ctx, task, reporter)
		results <- result
	}
}

// downloadWithRetry 带重试的下载（私有方法）
func (d *Downloader) downloadWithRetry(ctx context.Context, task DownloadTask, reporter ProgressReporter) DownloadResult {
	var lastErr error
	for attempt := 0; attempt <= d.opts.Retries; attempt++ {
		if attempt > 0 {
			// 重试前等待（指数退避）
			backoff := time.Duration(1<<uint(attempt-1)) * time.Second
			select {
			case <-ctx.Done():
				return DownloadResult{Task: task, Error: ctx.Err()}
			case <-time.After(backoff):
			}
		}

		result := d.downloadSingle(ctx, task, reporter)
		if result.Error == nil {
			return result
		}
		lastErr = result.Error
	}
	return DownloadResult{Task: task, Error: fmt.Errorf("重试 %d 次后仍失败: %w", d.opts.Retries, lastErr)}
}

// downloadSingle 单次下载尝试（私有方法）
func (d *Downloader) downloadSingle(ctx context.Context, task DownloadTask, reporter ProgressReporter) DownloadResult {
	task.StartTime = time.Now()

	// 创建 HTTP 请求
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, task.URL, nil)
	if err != nil {
		return DownloadResult{Task: task, Error: fmt.Errorf("创建请求失败: %w", err)}
	}

	// 发送请求
	resp, err := d.client.Do(req)
	if err != nil {
		return DownloadResult{Task: task, Error: fmt.Errorf("请求失败: %w", err)}
	}
	defer resp.Body.Close() // ← defer 经典用法：紧挨着资源创建

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return DownloadResult{Task: task, Error: fmt.Errorf("HTTP %d", resp.StatusCode)}
	}

	task.Size = resp.ContentLength

	// 创建输出文件
	filePath := task.FilePath
	if filePath == "" {
		filePath = filepath.Join(d.opts.OutputDir, filepath.Base(task.URL))
		if filePath == "." || filePath == "/" || filePath == d.opts.OutputDir {
			filePath = filepath.Join(d.opts.OutputDir, "download")
		}
	}

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return DownloadResult{Task: task, Error: fmt.Errorf("创建目录失败: %w", err)}
	}

	f, err := os.Create(filePath)
	if err != nil {
		return DownloadResult{Task: task, Error: fmt.Errorf("创建文件失败: %w", err)}
	}
	defer f.Close()

	// 带进度报告的拷贝
	downloaded, err := copyWithProgress(ctx, f, resp.Body, task.URL, task.Size, reporter)
	if err != nil {
		return DownloadResult{Task: task, Error: fmt.Errorf("下载失败: %w", err)}
	}

	task.EndTime = time.Now()
	task.FilePath = filePath

	// 计算平均速度
	duration := task.EndTime.Sub(task.StartTime).Seconds()
	avgSpeed := int64(0)
	if duration > 0 {
		avgSpeed = int64(float64(downloaded) / duration)
	}

	return DownloadResult{
		Task:        task,
		Downloaded:  downloaded,
		AvgSpeedBps: avgSpeed,
	}
}

// copyWithProgress 带进度报告的 io.Copy
// WHAT: 包装 io.Copy，每次读取后报告进度
// WHY: 这是 Go 的常见模式 —— 包装接口满足新需求
func copyWithProgress(ctx context.Context, dst io.Writer, src io.Reader, url string, total int64, reporter ProgressReporter) (int64, error) {
	buf := make([]byte, 32*1024) // 32KB 缓冲区
	var downloaded int64
	lastReportTime := time.Now()
	var lastDownloaded int64

	for {
		// 检查上下文
		select {
		case <-ctx.Done():
			return downloaded, ctx.Err()
		default:
		}

		n, readErr := io.ReadFull(src, buf)
		if n > 0 {
			if _, writeErr := dst.Write(buf[:n]); writeErr != nil {
				return downloaded, writeErr
			}
			downloaded += int64(n)

			// 每 200ms 或完成时报告进度（避免过于频繁）
			now := time.Now()
			if reporter != nil && (now.Sub(lastReportTime) > 200*time.Millisecond || readErr == io.ErrUnexpectedEOF || readErr == io.EOF) {
				elapsed := now.Sub(lastReportTime).Seconds()
				speed := int64(0)
				if elapsed > 0 {
					speed = int64(float64(downloaded-lastDownloaded) / elapsed)
				}
				reporter.OnProgress(url, downloaded, total, speed)
				lastReportTime = now
				lastDownloaded = downloaded
			}
		}

		if readErr == io.EOF || readErr == io.ErrUnexpectedEOF {
			return downloaded, nil
		}
		if readErr != nil {
			return downloaded, readErr
		}
	}
}
