// WHAT: 并发下载器入口 — 参数解析，任务编排
// WHY: 简洁的 main 函数，复杂逻辑交给 downloader 包
// 用法：go run main.go -urls urls.txt

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/user/go-basic-learn/level-11/downloader"
)

func main() {
	// ============================================================
	// 参数解析
	// ============================================================
	concurrency := flag.Int("concurrency", 3, "并发下载数")
	outputDir := flag.String("o", "./downloads", "输出目录")
	timeout := flag.Int("timeout", 30, "单个下载超时（秒）")
	retries := flag.Int("retries", 2, "失败重试次数")
	urlFlag := flag.String("url", "", "单个下载 URL")
	urlsFile := flag.String("urls", "", "包含 URL 列表的文件（每行一个）")
	flag.Parse()

	// 收集 URL
	var urls []string
	if *urlFlag != "" {
		urls = append(urls, *urlFlag)
	}
	if *urlsFile != "" {
		fileURLs, err := readURLsFromFile(*urlsFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "读取 URL 文件失败: %v\n", err)
			os.Exit(1)
		}
		urls = append(urls, fileURLs...)
	}
	if len(urls) == 0 {
		fmt.Fprintf(os.Stderr, "用法: %s -url <URL> 或 -urls <文件>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	// ============================================================
	// 构建下载任务
	// ============================================================
	tasks := make([]downloader.DownloadTask, 0, len(urls))
	for _, url := range urls {
		url = strings.TrimSpace(url)
		if url == "" || strings.HasPrefix(url, "#") {
			continue // 跳过空行和注释行
		}
		tasks = append(tasks, downloader.DownloadTask{
			URL: url,
		})
	}

	fmt.Printf("🚀 并发下载器启动\n")
	fmt.Printf("   任务数: %d | 并发数: %d | 输出: %s\n\n",
		len(tasks), *concurrency, *outputDir)

	// ============================================================
	// 信号处理 — 优雅取消
	// ============================================================
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\n\n⚠️  收到中断信号，正在取消下载...")
		cancel()
	}()

	// ============================================================
	// 执行下载
	// ============================================================
	opts := downloader.Options{
		Concurrency: *concurrency,
		OutputDir:   *outputDir,
		Timeout:     time.Duration(*timeout) * time.Second,
		Retries:     *retries,
	}

	dl := downloader.NewDownloader(opts)
	progress := downloader.NewConsoleProgress()

	startTime := time.Now()
	results := dl.Download(ctx, tasks, progress)
	elapsed := time.Since(startTime)

	// ============================================================
	// 输出结果报告
	// ============================================================
	fmt.Println("\n\n========== 下载报告 ==========")

	var (
		successCount int
		failCount    int
		totalBytes   int64
	)
	for _, r := range results {
		fmt.Printf("  %s\n", r.String())
		if r.Error == nil {
			successCount++
			totalBytes += r.Downloaded
		} else {
			failCount++
		}
	}

	fmt.Println("\n-------------------------------")
	fmt.Printf("  总计: %d 成功, %d 失败\n", successCount, failCount)
	fmt.Printf("  下载量: %s\n", formatSizeStr(totalBytes))
	fmt.Printf("  总用时: %v\n", elapsed.Round(time.Millisecond))
	if elapsed.Seconds() > 0 {
		fmt.Printf("  平均速度: %s/s\n", formatSizeStr(int64(float64(totalBytes)/elapsed.Seconds())))
	}
	fmt.Println("===============================")

	if failCount > 0 {
		os.Exit(1)
	}
}

// readURLsFromFile 从文件读取 URL 列表（每行一个）
func readURLsFromFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer f.Close()

	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls, scanner.Err()
}

// formatSizeStr 格式化文件大小（辅助函数）
func formatSizeStr(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
