// WHAT: Bug 01 — goroutine 闭包变量捕获（Go 最经典的并发 Bug）
// ERROR: goroutine 闭包捕获循环变量时，所有 goroutine 看到同一个变量
//        Go 1.22 之前这是最常见的并发 Bug
//
// ============================================================
// 运行结果 (Go 1.21-):
// 3 3 3 （全是最后一个值）
//
// 运行结果 (Go 1.22+):
// 0 1 2 / 1 0 2 / ... （随机，但都是不同值）
// ============================================================
//
// 为什么会这样：
//   在 Go 1.21 之前，for 循环中 i 的地址不变，只是值变化。
//   闭包捕获的是变量地址（不是值），所以所有 goroutine 读到的
//   是循环结束后的 i（最后一次的值）。
//
//   Go 1.22 修复了：每次迭代创建新变量。
//
// CONTRAST（与已知语言对比）：
//   - Rust: for i in 0..3 { spawn(move || println!("{}", i)) }
//           move 关键字强制捕获值 —— 编译器提醒你
//   - TS:   for (let i = 0; i < 3; i++) { setTimeout(() => console.log(i)) }
//           let 是块作用域，每次迭代新变量 —— 安全
//   - Go 1.21-: 循环变量重用 —— 陷阱！
//   - Go 1.22+: 修复了
//
// 如何修复（兼容所有版本）：
//   循环内拷贝：i := i 或传参：go func(id int) {}

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("--- goroutine 闭包变量捕获 ---")

	var wg sync.WaitGroup

	fmt.Print("Go 1.21- bug 版本 (类似): ")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%d ", i) // 闭包捕获 i 的地址（Go 1.21 前都是同一个）
		}()
	}
	wg.Wait()
	fmt.Println("(Go 1.22+ 已经安全)\n")

	// 兼容所有版本的写法
	fmt.Print("兼容写法（传参）: ")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("%d ", id)
		}(i)
	}
	wg.Wait()
	fmt.Println()

	fmt.Print("兼容写法（拷贝）: ")
	for i := 0; i < 3; i++ {
		i := i // ← 这行看起来像废话，但创建了新变量
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%d ", i)
		}()
	}
	wg.Wait()
	fmt.Println()
}
