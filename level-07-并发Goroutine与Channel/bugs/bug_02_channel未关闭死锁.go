// WHAT: Bug 02 — Channel 未关闭导致死锁
// ERROR: range channel 时，如果 channel 未关闭且没有新数据，
//        range 会一直阻塞 → 所有 goroutine 睡眠 → deadlock!
//
// ============================================================
// 运行结果：
// fatal error: all goroutines are asleep - deadlock!
// ============================================================
//
// 为什么会这样：
//   range channel 在没有数据且 channel 未关闭时，会阻塞等待。
//   如果没有任何 goroutine 会再往 channel 发送数据，Go runtime
//   检测到所有 goroutine 都在睡眠 → fatal deadlock
//
// CONTRAST（与已知语言对比）：
//   - Rust: tokio mpsc 的 Receiver 在 Sender 全部 drop 后返回 None
//   - Go:   channel 不会"知道"发送方已全部退出，必须显式 close
//           或者确保你知道什么时候不再发送
//
// 如何修复：
//   1. 发送方 close(ch) 后，接收方 range 自动退出
//   2. 用 select + case <-done: 检测终止信号

package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	// 发送 3 个值
	ch <- 1
	ch <- 2
	ch <- 3

	// BUG：如果这里没有 close(ch)，range 会一直等待第 4 个值
	// 而没有任何 goroutine 会发送第 4 个值 → deadlock!

	// 取消注释下面这行，修复死锁：
	// close(ch)

	fmt.Println("开始 range channel（如果没有 close，下面会死锁）")

	// 以下 range 在 channel 未关闭时会永久阻塞
	for v := range ch {
		fmt.Println("收到:", v)
		// 读完 3 个值后，range 继续等 → 死锁
		if v == 3 {
			// 修复：读完后关闭或者不用 range
			break
		}
	}

	fmt.Println("注意：如果看到这行说明没有死锁（break 提前退出了）")
	fmt.Println("    但如果删掉 break，用 range 读完就会 deadlock")
	fmt.Println("    修复：读完记得 close(ch)")
}
