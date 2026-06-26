// WHAT: Level 07 主代码 — Goroutine 与 Channel
// WHY: Go 的并发模型是其最大的卖点。理解 CSP 模型 vs async/await vs 线程
// CONTRAST: TS 的 async/await（单线程），Rust 的 async + 线程，Go 的 goroutine

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("========== Goroutine 与 Channel ==========\n")

	// ============================================================
	// 第一部分：goroutine — 轻量级并发
	// ============================================================
	fmt.Println("--- 1. goroutine 基础 ---")
	// WHAT: go f() 启动一个新的 goroutine，与主 goroutine 并发执行
	// WHY: goroutine 不是线程！它是 Go 运行时管理的用户态"协程"
	//      一个 goroutine 初始栈只有 ~2KB（OS 线程通常 1MB+）
	//      可以同时运行数百万个 goroutine
	// CONTRAST:
	//   - TS:   async function → 事件循环的微任务（单线程）
	//   - Rust: tokio::spawn → 异步任务（协作式调度）
	//   - Go:   go func() → goroutine（抢占式调度，GMP）

	go func() {
		fmt.Println("  goroutine: 你好，来自另一个 goroutine!")
	}()
	time.Sleep(10 * time.Millisecond) // 给 goroutine 执行机会（实际用 WaitGroup）

	// ============================================================
	// 第二部分：sync.WaitGroup — 等待所有 goroutine
	// ============================================================
	fmt.Println("\n--- 2. WaitGroup ---")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 计数器 +1
		go func(id int) {
			defer wg.Done() // 完成后计数器 -1
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Printf("  Worker %d 完成\n", id)
		}(i)
	}
	wg.Wait() // 阻塞直到计数器归零
	fmt.Println("  所有 Worker 完成\n")

	// ============================================================
	// 第三部分：Channel — Goroutine 之间的管道
	// ============================================================
	fmt.Println("--- 3. Channel ---")
	// WHAT: channel 是 goroutine 之间通信的管道
	// WHY: Go 的哲学："不要用共享内存来通信，用通信来共享内存"
	//      Don't communicate by sharing memory; share memory by communicating.
	// CONTRAST:
	//   - TS:   没有等效物（单线程），或 Worker.postMessage
	//   - Rust: tokio::sync::mpsc 最接近

	// 无缓冲 channel — 发送方阻塞直到接收方就绪（同步）
	unbuffered := make(chan string)
	go func() {
		unbuffered <- "ping" // 发送，阻塞直到有人接收
	}()
	msg := <-unbuffered // 接收
	fmt.Printf("  无缓冲 channel: %s\n", msg)

	// 有缓冲 channel — 缓冲满之前不阻塞（异步）
	buffered := make(chan int, 3)
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Printf("  有缓冲 channel (cap=3): len=%d, cap=%d\n", len(buffered), cap(buffered))
	// buffered <- 4 // ← 会阻塞！缓冲满了

	// 关闭 channel
	close(buffered)
	// 从已关闭的 channel 读取：先读缓冲中的数据，然后零值
	v, ok := <-buffered
	fmt.Printf("  从关闭的 channel 读: %d, ok=%t\n", v, ok)
	v, ok = <-buffered
	fmt.Printf("  继续读: %d, ok=%t\n", v, ok)
	// 等等...一直读到缓冲区空，然后每次都是零值
	fmt.Println()

	// ============================================================
	// 第四部分：Producer-Consumer 模式
	// ============================================================
	fmt.Println("--- 4. Producer-Consumer ---")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// 启动 2 个 worker
	var workerWg sync.WaitGroup
	for w := 1; w <= 2; w++ {
		workerWg.Add(1)
		go worker(w, jobs, results, &workerWg)
	}

	// 发送 5 个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // 关闭 jobs → workers 的 range 会退出

	// 等待所有 worker 结束
	workerWg.Wait()
	close(results) // 关闭 results

	// 收集结果
	fmt.Print("  结果: ")
	for r := range results {
		fmt.Printf("%d ", r)
	}
	fmt.Println("\n")

	// ============================================================
	// 第五部分：select — 多路复用
	// ============================================================
	fmt.Println("--- 5. select 多路复用 ---")
	// WHAT: select 类似 Unix select()，同时监听多个 channel
	//       哪个 channel 先就绪就执行哪个 case
	// WHY: 实现超时、非阻塞、多通道选择等。
	// CONTRAST:
	//   - TS:   Promise.race() / Promise.any() 类似
	//   - Rust: tokio::select! 宏

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "来自 ch1"
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		ch2 <- "来自 ch2 (先到!)"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Printf("  %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("  %s\n", msg2)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  超时!")
	}

	// ============================================================
	// 第六部分：select 超时模式
	// ============================================================
	fmt.Println("\n--- 6. 超时模式 ---")
	slowCh := make(chan string)
	select {
	case msg := <-slowCh:
		fmt.Printf("  收到: %s\n", msg)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("  操作超时（50ms）")
	}

	// ============================================================
	// 第七部分：range channel（迭代 channel）
	// ============================================================
	fmt.Println("\n--- 7. range channel ---")
	stream := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			stream <- i
		}
		close(stream) // 必须关闭，否则 range 永不退出
	}()
	fmt.Print("  range channel: ")
	for v := range stream {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// ============================================================
	// 第八部分：Pipeline 模式
	// ============================================================
	fmt.Println("\n--- 8. Pipeline 模式 ---")
	// 经典的 Go 并发模式：gen → sq → print
	for n := range sq(gen(1, 2, 3, 4, 5)) {
		fmt.Printf("  %d ", n)
	}
	fmt.Println()

	fmt.Println("\n✅ Level 07 完成！")
}

// worker 消费者（从 jobs 读取，写入 results）
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("  Worker %d 处理任务 %d\n", id, j)
		time.Sleep(20 * time.Millisecond) // 模拟工作
		results <- j * 2
	}
}

// gen 生成器 —— 返回只读 channel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq 平方 —— 接收只读，返回只读
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
