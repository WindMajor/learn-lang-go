// WHAT: 进度追踪器实现
// WHY: 独立的关注点，终端 UI 和下载逻辑完全解耦

package downloader

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// ConsoleProgress 终端进度报告器（实现 ProgressReporter 接口）
// WHAT: 首字母大写 → 导出，其他包可以使用
type ConsoleProgress struct {
	mu          sync.Mutex
	tasks       map[string]*taskProgress
	startTime   time.Time
}

// taskProgress 单个任务的进度（私有，仅包内可见）
type taskProgress struct {
	downloaded int64
	total      int64
	speedBps   int64
	done       bool
}

// NewConsoleProgress 创建终端进度报告器（构造函数模式）
func NewConsoleProgress() *ConsoleProgress {
	return &ConsoleProgress{
		tasks:     make(map[string]*taskProgress),
		startTime: time.Now(),
	}
}

// OnProgress 实现 ProgressReporter 接口
func (cp *ConsoleProgress) OnProgress(url string, downloaded, total, speedBps int64) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if _, exists := cp.tasks[url]; !exists {
		cp.tasks[url] = &taskProgress{}
	}
	tp := cp.tasks[url]
	tp.downloaded = downloaded
	tp.total = total
	tp.speedBps = speedBps

	if total > 0 && downloaded >= total {
		tp.done = true
	}

	cp.render()
}

// render 渲染进度条（私有方法）
func (cp *ConsoleProgress) render() {
	// 使用 \r 实现行内更新
	fmt.Print("\r\033[K") // 清除当前行

	var parts []string
	for url, tp := range cp.tasks {
		// 截取 URL 末尾作为显示名
		displayName := shortenURL(url)
		if tp.done {
			parts = append(parts, fmt.Sprintf("%s ✅ %s",
				displayName, formatSize(tp.downloaded)))
		} else if tp.total > 0 {
			pct := float64(tp.downloaded) / float64(tp.total) * 100
			bar := progressBar(int(pct), 10)
			parts = append(parts, fmt.Sprintf("%s %s %.1f%% %s/s",
				displayName, bar, pct, formatSize(tp.speedBps)))
		} else {
			parts = append(parts, fmt.Sprintf("%s ↓ %s/s",
				displayName, formatSize(tp.speedBps)))
		}
	}
	fmt.Print(strings.Join(parts, " | "))
}

// shortenURL 缩短 URL 显示（私有函数）
func shortenURL(url string) string {
	// 取最后一个 / 之后的部分，或域名
	idx := strings.LastIndex(url, "/")
	if idx >= 0 && idx < len(url)-1 {
		return url[idx+1:]
	}
	return url
}

// progressBar 生成进度条（私有函数）
func progressBar(pct, width int) string {
	filled := pct * width / 100
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
	return fmt.Sprintf("[%s]", bar)
}
