// WHAT: Bug 01 — defer 参数在声明时求值（不是执行时！）
// ERROR: defer 语句的参数在 defer 声明时就求值了，而不是在函数返回时
//        这是 Go 中最容易踩的坑之一
//
// ============================================================
// 运行结果：
// 最终 x = 100
// defer x = 1  ← 不是 100！
// ============================================================
//
// 为什么会这样：
//   defer fmt.Println(x) —— fmt.Println 的参数 x 在 defer 声明时
//   （即这行代码执行时）就被求值了。之后 x 被修改为 100，但 defer
//   的参数已经是旧值了。
//
// 这就像 Rust 中你 move 了一个值到闭包中：
//   let x = 1;
//   let deferred = || println!("{}", x); // x 被捕获
//   x = 100; // 不影响 deferred 中已经捕获的 x（如果是 move）
//
// CONTRAST（与已知语言对比）：
//   - Rust: 闭包捕获变量，取决于 Fn/FnMut/FnOnce（move 或借引用）
//   - TS:   闭包捕获变量引用（总是看到最新值）
//   - Go:   defer 参数在声明时求值！如果想看到最新值，用闭包包装
//
// 如何修复：
//   方法 1：用闭包包装（闭包捕获的是变量引用，看到最新值）
//     defer func() { fmt.Println(x) }()
//   方法 2：传指针（变通方案，不推荐）
//     defer func(p *int) { fmt.Println(*p) }(&x)

package main

import "fmt"

func main() {
	// BUG 演示：defer 参数声明时求值
	x := 1
	defer fmt.Println("defer x =", x) // ← x 在这里被求值！值是 1
	// 等价于: defer fmt.Println("defer x =", 1)
	x = 100
	fmt.Println("最终 x =", x)
	// 输出：最终 x = 100  →  defer x = 1  (不是 100！)

	// ============================================================
	// 修复版本
	// ============================================================
	fmt.Println("\n--- 修复演示 ---")
	y := 1
	defer func() {
		fmt.Println("defer y (闭包) =", y) // 闭包捕获变量引用，看到最新值
	}()
	y = 100
	fmt.Println("最终 y =", y)
	// 输出：最终 y = 100  →  defer y (闭包) = 100 ✅
}
