// WHAT: 下载器单元测试
// WHY: 验证核心逻辑的正确性，遵循 Go 的表驱动测试范式

package downloader

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestFormatSize 测试文件大小格式化
func TestFormatSize(t *testing.T) {
	tests := []struct {
		bytes int64
		want  string
	}{
		{0, "0 B"},
		{500, "500 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{1048576, "1.0 MB"},
		{1073741824, "1.0 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := formatSize(tt.bytes)
			if got != tt.want {
				t.Errorf("formatSize(%d) = %q, want %q", tt.bytes, got, tt.want)
			}
		})
	}
}

// TestDownloadSingle 测试单个文件下载
func TestDownloadSingle(t *testing.T) {
	// 创建测试 HTTP 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// 写入 1KB 测试数据
		data := make([]byte, 1024)
		for i := range data {
			data[i] = byte(i % 256)
		}
		w.Write(data)
	}))
	defer server.Close()

	// 创建临时输出目录
	tmpDir := t.TempDir()

	// 创建下载器
	opts := DefaultOptions()
	opts.OutputDir = tmpDir
	opts.Concurrency = 1
	dl := NewDownloader(opts)

	// 构建任务
	task := DownloadTask{
		URL:      server.URL + "/test.bin",
		FilePath: filepath.Join(tmpDir, "test.bin"),
	}

	// 执行下载
	ctx := context.Background()
	results := dl.Download(ctx, []DownloadTask{task}, nil)

	// 验证结果
	if len(results) != 1 {
		t.Fatalf("期望 1 个结果，得到 %d", len(results))
	}

	result := results[0]
	if result.Error != nil {
		t.Fatalf("下载失败: %v", result.Error)
	}
	if result.Downloaded != 1024 {
		t.Errorf("期望下载 1024 字节，实际 %d", result.Downloaded)
	}

	// 验证文件存在且大小正确
	stat, err := os.Stat(result.Task.FilePath)
	if err != nil {
		t.Fatalf("文件不存在: %v", err)
	}
	if stat.Size() != 1024 {
		t.Errorf("文件大小 %d，期望 1024", stat.Size())
	}
}

// TestDownloadWithCancellation 测试取消下载
func TestDownloadWithCancellation(t *testing.T) {
	// 创建慢速服务器（模拟大文件）
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// 每写 1 字节就等一会，模拟慢速下载
		for i := 0; i < 100; i++ {
			select {
			case <-r.Context().Done():
				return
			default:
				w.Write([]byte{byte(i)})
				time.Sleep(10 * time.Millisecond)
			}
		}
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	opts := DefaultOptions()
	opts.OutputDir = tmpDir
	dl := NewDownloader(opts)

	task := DownloadTask{
		URL:      server.URL + "/slow.bin",
		FilePath: filepath.Join(tmpDir, "slow.bin"),
	}

	// 创建可取消的 context
	ctx, cancel := context.WithCancel(context.Background())
	// 100ms 后取消
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	results := dl.Download(ctx, []DownloadTask{task}, nil)

	if len(results) != 1 {
		t.Fatalf("期望 1 个结果，得到 %d", len(results))
	}
	if results[0].Error == nil {
		t.Error("期望下载被取消，但没有错误")
	}
}

// TestResultString 测试结果字符串格式化
func TestResultString(t *testing.T) {
	now := time.Now()
	result := DownloadResult{
		Task: DownloadTask{
			URL:       "http://example.com/file.zip",
			FilePath:  "/tmp/file.zip",
			StartTime: now,
			EndTime:   now.Add(1 * time.Second),
		},
		Downloaded:  1048576, // 1MB
		AvgSpeedBps: 1048576,
	}

	str := result.String()
	fmt.Println("结果:", str)
	if str == "" {
		t.Error("结果字符串不应为空")
	}
}

// TestDefaultOptions 测试默认配置
func TestDefaultOptions(t *testing.T) {
	opts := DefaultOptions()
	if opts.Concurrency != 3 {
		t.Errorf("默认并发数应为 3，实际 %d", opts.Concurrency)
	}
	if opts.OutputDir != "./downloads" {
		t.Errorf("默认输出目录应为 ./downloads，实际 %s", opts.OutputDir)
	}
	if opts.Retries != 2 {
		t.Errorf("默认重试次数应为 2，实际 %d", opts.Retries)
	}
}

// TestHTTPErrorHandling 测试 HTTP 错误处理
func TestHTTPErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	tmpDir := t.TempDir()
	opts := DefaultOptions()
	opts.OutputDir = tmpDir
	dl := NewDownloader(opts)

	task := DownloadTask{
		URL:      server.URL + "/notfound",
		FilePath: filepath.Join(tmpDir, "notfound"),
	}

	ctx := context.Background()
	results := dl.Download(ctx, []DownloadTask{task}, nil)

	if results[0].Error == nil {
		t.Error("期望 HTTP 404 错误，但没有")
	}
}

// BenchmarkDownloadSpeed 基准测试（跳过实际网络请求）
func BenchmarkFormatSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatSize(1234567890)
	}
}
