// WHAT: Bug 01 — 包级变量不能用 := 短声明
// ERROR: := 只能在函数内使用，包级别必须用 var
//
// ============================================================
// 编译错误信息：
// syntax error: non-declaration statement outside function body
// ============================================================
//
// 为什么会这样：
//   Go 的语法规定：:= 是"短变量声明"，只能出现在函数体内。
//   包级别（全局作用域）只能使用 var、const、type、func 关键字。
//   这是有意为之 —— 包级别的声明应该显式、可预测。
//
// CONTRAST（与已知语言对比）：
//   - Rust: let 也只能在函数内，全局用 static/const
//   - TS:   顶层也可以用 const/let（模块作用域）
//   - C:    全局变量必须显式指定类型或用 auto（C23）
//   - Go:   包级别不能用 :=
//
// 如何修复：
//   把 := 改成 var（带或不带类型）：
//     var maxConnections = 100
//     或
//     var maxConnections int = 100
//
// ============================================================

package main

import "fmt"

// BUG: 包级别不能用 :=
// maxConnections := 100  // ← 取消注释：syntax error

// 正确写法
var maxConnections = 100

func main() {
	fmt.Printf("maxConnections = %d\n", maxConnections)
	// 注意：这个 bug 文件实际上能编译通过（因为我们写了正确的版本）
	// 请取消上面那个 := 行，注释掉 var 行，然后 go build 看错误
}
