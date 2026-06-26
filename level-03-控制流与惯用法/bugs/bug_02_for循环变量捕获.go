// WHAT: Bug 02 — for 循环变量捕获陷阱（Go 1.22 之前经典 Bug）
// ERROR: 在循环中启动 goroutine 或创建闭包时，循环变量在每次迭代被重用
//        导致所有闭包捕获的都是同一个变量（最终值）
//
// ============================================================
// 运行结果（Go 1.21 及之前）：
// 3 3 3 （全部是最后一个值！）
//
// 运行结果（Go 1.22+）：
// 0 1 2 （已修复）
// ============================================================
//
// 为什么会这样：
//   Go 1.21 及之前：for 循环中 i 的地址不变，每次迭代只改值
//   Go 1.22+：每次迭代创建新变量（类似 Rust 的 move 语义）
//
// CONTRAST（与已知语言对比）：
//   - Rust: for 循环中迭代器 move，不会有此问题
//   - TS:   let 是块作用域，每次迭代新变量（for...of 也一样）
//   - Go 1.21: 循环变量重用！最经典的 Go 陷阱之一
//   - Go 1.22: 修复了，和 TS 行为一致
//
// 如何修复（所有版本通用）：
//   在循环体内拷贝：i := i（利用作用域创建新变量）

package main

import "fmt"

func main() {
	fmt.Println("Go 版本:", "1.22+ 已修复此 Bug，但旧版本需要注意")

	// 模拟闭包捕获
	var funcs []func() int

	for i := 0; i < 3; i++ {
		// Go 1.22+：每次迭代 i 是新变量，安全
		// Go 1.21 及之前：i 是同一个变量，捕获的是最终值
		funcs = append(funcs, func() int { return i })
	}

	for _, f := range funcs {
		fmt.Printf("%d ", f())
	}
	fmt.Println()
	fmt.Println("Go 1.22+ 输出: 0 1 2 ✅")
	fmt.Println("Go 1.21- 输出: 3 3 3 ❌")

	// 兼容所有版本的写法：
	fmt.Print("兼容写法: ")
	funcs2 := []func() int{}
	for i := 0; i < 3; i++ {
		i := i // ← 这行看起来多余，但是所有版本都安全！
		funcs2 = append(funcs2, func() int { return i })
	}
	for _, f := range funcs2 {
		fmt.Printf("%d ", f())
	}
	fmt.Println()
}
