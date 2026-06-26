// Package downloader 实现一个并发文件下载器
// WHAT: 定义下载器的核心类型和接口
// WHY: 类型定义集中管理，遵循 Go 的"一包一职责"原则
// CONTRAST:
//   - Rust: 类似的模块拆分，但 Rust 用 mod + pub 控制可见性
//   - TS:   通常一个文件导出多个 interface/type
//   - Go:   按概念拆分文件，同一包内类型自动共享

package downloader

import (
	"fmt"
	"time"
)

// DownloadTask 表示一个下载任务
type DownloadTask struct {
	// URL 下载地址（导出）
	URL string
	// FilePath 保存路径（导出）
	FilePath string
	// Size 文件大小（字节），下载前可能为 0
	Size int64
	// StartTime 下载开始时间
	StartTime time.Time
	// EndTime 下载完成时间
	EndTime time.Time
}

// DownloadResult 下载结果
type DownloadResult struct {
	Task        DownloadTask
	Error       error
	Downloaded  int64  // 实际下载字节数
	AvgSpeedBps int64  // 平均速度（字节/秒）
}

// Status 返回结果状态字符串（导出方法）
func (r DownloadResult) Status() string {
	if r.Error != nil {
		return "失败"
	}
	return "成功"
}

// String 实现 fmt.Stringer 接口（导出方法）
func (r DownloadResult) String() string {
	if r.Error != nil {
		return fmt.Sprintf("[%s] %s → 失败: %v", r.Status(), r.Task.URL, r.Error)
	}
	duration := r.Task.EndTime.Sub(r.Task.StartTime)
	speed := formatSize(r.AvgSpeedBps) + "/s"
	return fmt.Sprintf("[%s] %s → %s (用时 %v, 速度 %s)",
		r.Status(), r.Task.URL, formatSize(r.Downloaded), duration.Round(time.Millisecond), speed)
}

// ProgressReporter 进度报告接口（隐式接口）
// WHAT: 任何能 OnProgress 的类型都是进度报告器
//       你可以实现为终端输出、WebSocket 推送、日志写入等
// WHY: 接口解耦下载逻辑和进度展示 —— Go 的设计哲学
// CONTRAST:
//   - TS: 类似的 interface ProgressReporter { onProgress(...) }
//   - Rust: trait ProgressReporter { fn on_progress(&self, ...) }
//   - Go: 隐式实现，无需 implements/impling
type ProgressReporter interface {
	// OnProgress 收到进度更新
	// url: 正在下载的 URL
	// downloaded: 已下载字节数
	// total: 总字节数（未知时为 0）
	// speedBps: 当前速度
	OnProgress(url string, downloaded, total, speedBps int64)
}

// Options 下载器配置选项
type Options struct {
	// Concurrency 并发下载数（导出）
	Concurrency int
	// OutputDir 输出目录
	OutputDir string
	// Timeout 单个下载超时时间
	Timeout time.Duration
	// Retries 失败重试次数
	Retries int
}

// DefaultOptions 返回默认配置（导出函数）
func DefaultOptions() Options {
	return Options{
		Concurrency: 3,
		OutputDir:   "./downloads",
		Timeout:     30 * time.Second,
		Retries:     2,
	}
}

// formatSize 格式化文件大小（私有辅助函数）
func formatSize(bytes int64) string {
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
