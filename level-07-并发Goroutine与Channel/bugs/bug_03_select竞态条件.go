// WHAT: Bug 03 — Select 竞态条件与 channel 泄漏
// ERROR: select 的 case 是随机选择的（多个 case 同时就绪时）
//        这可能导致竞态条件，且未选择的 case 中的 goroutine 可能泄露
//
// ============================================================
// 运行结果（随机）：
// 可能收到: msg1 或 msg2
// ============================================================
//
// 为什么会这样：
//   Go 的 select 在多个 case 同时就绪时，**随机选择**一个执行。
//   这是刻意为之 —— 防止某个 case 被饿死。
//   但这意味着你不能依赖 select 的优先级。
//
//   另外，如果 select 选择了某个 case，其他 case 的 channel 操作
//   不会被执行，相关 goroutine 可能永远阻塞（goroutine 泄漏）。
//
// CONTRAST（与已知语言对比）：
//   - Rust: tokio::select! 也是随机（或 biased 模式有优先级）
//   - TS:   Promise.race() —— 第一个完成的胜出

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan struct{})

	// goroutine 1：立即发送
	go func() {
		ch1 <- "msg1"
	}()

	// goroutine 2：也立即发送
	go func() {
		ch2 <- "msg2"
	}()

	// BUG：两个 case 同时就绪，select 随机选一个
	// 未被选择的 case 的 goroutine 会一直阻塞
	// 但如果 ch1/ch2 是无缓冲的，发送 goroutine 会阻塞在那里 →

	// 修复：两个都接收
	time.Sleep(10 * time.Millisecond)
	go func() {
		// 这个 goroutine 接收剩余的那个
		select {
		case msg := <-ch1:
			fmt.Println("回收 ch1:", msg)
		case msg := <-ch2:
			fmt.Println("回收 ch2:", msg)
		}
		close(done)
	}()

	time.Sleep(10 * time.Millisecond)
	select {
	case msg := <-ch1:
		fmt.Println("选中 ch1:", msg)
	case msg := <-ch2:
		fmt.Println("选中 ch2:", msg)
	}

	<-done // 等待回收完成
	fmt.Println("注意：select 随机性意味着不要依赖优先级")
	fmt.Println("    未被选中的 case 中的 goroutine 可能泄漏")
}
