// WHAT: Bug 01 — 接口 nil 值陷阱（Go 最著名的陷阱之一！）
// ERROR: 接口值由 (type, value) 组成。当 value 是 nil 但 type 不是 nil 时，
//        接口整体不为 nil！这导致 nil 检查失效。
//
// ============================================================
// 运行结果：
// g 不是 nil！ (g == nil: false)
// ============================================================
//
// 为什么会这样：
//   Go 的接口内部是 (类型, 值) 的二元组。
//   只有 type 和 value 都为 nil 时，接口才为 nil。
//   当返回一个具体类型 *T（即使值为 nil），赋给接口时：
//     var g Greeter = (*CustomGreeter)(nil)  →  接口 != nil！
//     因为接口的类型部分 = *CustomGreeter（不为 nil）
//
//   这是 Go 面试最高频的问题：函数返回 *MyError(nil)，外面 err != nil 判断失效。
//
// CONTRAST（与已知语言对比）：
//   - TS:   `const g: Greeter | null = null;` — null 判断安全
//   - Rust: `let g: Option<&dyn Greeter> = None;` — Option 强制处理 None
//   - Go:   接口的 nil 判断陷阱 —— 最隐蔽也最常见的错误来源
//
// 如何修复：
//   1. 函数直接返回 nil 接口，不返回 nil 具体类型指针
//   2. Go 1.22+ 用闭包或特殊处理避免返回 nil 指针作为接口

package main

import "fmt"

// Greeter 接口
type Greeter interface {
	Greet() string
}

// Person 结构体
type Person struct {
	Name string
}

func (p *Person) Greet() string {
	if p == nil {
		return "nil person greeting"
	}
	return "Hello, " + p.Name
}

// BUG: 返回 nil 指针作为接口！
func getGreeterBug() Greeter {
	var p *Person // p == nil
	// return p ← 如果取消注释，返回的是类型=*Person, 值=nil → 接口 != nil！
	return nil // 正确：直接返回 nil 接口
}

// 修复：当你需要返回"可能为空"的接口时
func getGreeterFixed(name string) Greeter {
	if name == "" {
		return nil // 直接返回 nil 接口，不要返回 nil 指针
	}
	return &Person{Name: name}
}

func main() {
	// 演示 1：nil 指针赋给接口
	var p *Person = nil
	var g Greeter = p

	fmt.Printf("p == nil: %t\n", p == nil)  // true
	fmt.Printf("g == nil: %t\n", g == nil)  // false！← 关键！
	fmt.Printf("g.Greet(): %s\n\n", g.Greet()) // 但方法可以调用（如果方法处理了 nil）

	// 这是因为接口 g 的内部表示是：
	//   (type=*Person, value=nil)
	// type 不为 nil，所以 g != nil

	fmt.Printf("getGreeterBug(): %v == nil? %t\n", getGreeterBug(), getGreeterBug() == nil)
	fmt.Printf("getGreeterFixed(\"\"): %v == nil? %t\n", getGreeterFixed(""), getGreeterFixed("") == nil)
	fmt.Printf("getGreeterFixed(\"Alice\"): %v == nil? %t\n", getGreeterFixed("Alice"), getGreeterFixed("Alice") == nil)
}
