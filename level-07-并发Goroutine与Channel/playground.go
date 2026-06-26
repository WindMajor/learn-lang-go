// WHAT: 沙盒文件 — goroutine + channel 实验
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("🎮 Level 07 沙盒\n")

	// ======== 简单 goroutine ========
	fmt.Println("--- goroutine ---")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("goroutine %d\n", n)
		}(i)
	}
	wg.Wait()

	// ======== channel 通信 ========
	fmt.Println("--- channel ---")
	ch := make(chan string, 1)
	ch <- "message"
	fmt.Println(<-ch)

	// ======== select 实验 ========
	fmt.Println("--- select ---")
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "第一个"
	select {
	case msg := <-ch1:
		fmt.Println("收到:", msg)
	case msg := <-ch2:
		fmt.Println("收到:", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("超时")
	}

	// ======== 你的实验区 ========
	fmt.Println("\n--- 你的实验 ---")

	fmt.Println("✅ 沙盒完成！")
}
