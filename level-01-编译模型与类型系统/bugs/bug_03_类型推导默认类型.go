// WHAT: Bug 03 — 类型推导的默认类型陷阱
// ERROR: 使用 := 推导类型时，整数默认推导为 int，浮点数默认 float64
//        当你需要 int32 或 float32 时，这种默认行为可能导致精度问题或类型不匹配
//
// ============================================================
// 编译错误信息（关键行）：
// cannot use x (variable of type int) as int32 value in argument to fn
// ============================================================
//
// 为什么会这样：
//   Go 的 := 类型推导规则：
//   - 整数字面量 → int（平台相关，64 位系统是 int64）
//   - 浮点字面量 → float64
//   - 复数字面量 → complex128
//   - 字符串字面量 → string
//   - 布尔字面量 → bool
//   - 字符字面量 → rune（int32 别名）
//
//   int 和 int32 在 Go 中是**完全不同的类型**，需要显式转换。
//   不像 C 中 int 和 long 可以自动提升转换。
//
// CONTRAST（与已知语言对比）：
//   - Rust: let x = 42;  // 推导为 i32（固定），需要时再 as 转换
//   - TS:   let x = 42;  // number（IEEE 754 双精度），无所谓
//   - C:    auto x = 42; // C23 引入，推导为 int
//   - Go:   x := 42      // 推导为 int（平台相关，不固定！）
//
//   关键差异：Go 的 int 大小随平台变化，这可能导致跨平台 Bug！
//
// 如何修复：
//   需要特定类型时，不要用 :=，用 var 显式声明：
//     var x int32 = 42       // 明确指定 int32
//     var y float32 = 3.14   // 明确指定 float32
//   或者用类型转换：
//     x := int32(42)
//
// ============================================================

package main

import "fmt"

// WHAT: 这个函数明确要求 int32 类型参数
// WHY: 当你和 C 库交互或做二进制处理时，通常需要精确宽度类型
func needInt32(n int32) {
	fmt.Printf("needInt32 收到: %d (类型: %T)\n", n, n)
}

func needFloat32(f float32) {
	fmt.Printf("needFloat32 收到: %f (类型: %T)\n", f, f)
}

func main() {
	// BUG 1：:= 推导出来的整数是 int，不是 int32
	x := 42
	fmt.Printf("x := 42 → 类型是 %T (不是 int32!)\n\n", x)

	// needInt32(x) // ← 取消注释会报错：cannot use x (variable of type int) as int32 value

	// BUG 2：:= 推导出来的浮点是 float64，不是 float32
	y := 3.14
	fmt.Printf("y := 3.14 → 类型是 %T (不是 float32!)\n\n", y)

	// needFloat32(y) // ← 取消注释会报错：cannot use y (variable of type float64) as float32 value

	// BUG 3：字符字面量推导为 rune（int32），不是 byte
	c1 := 'A'  // rune!
	c2 := byte('A') // 这才是 byte
	fmt.Printf("c1 := 'A' → 类型是 %T\n", c1)
	fmt.Printf("c2 := byte('A') → 类型是 %T\n\n", c2)

	// ============================================================
	// 修复版本
	// ============================================================
	//
	// x32 := int32(42)       // 显式转换为 int32
	// needInt32(x32)        // 正常
	//
	// y32 := float32(3.14)  // 显式转换为 float32
	// needFloat32(y32)      // 正常
	//
	// var x32 var int32 = 42 // 或者用 var 声明时指定类型
	//
	// 或者更 Go 风格：var x int32 = 42

	// 正确的写法
	needInt32(int32(x))
	needFloat32(float32(y))

	fmt.Println("✅ 此文件是 bug 演示，展示了类型推导的默认类型问题")
}
