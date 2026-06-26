// WHAT: Bug 03 — 空接口类型断言 panic
// ERROR: 不带 comma-ok 的类型断言失败时会 panic
//        断言到不匹配的类型：panic: interface conversion
//
// ============================================================
// panic: interface conversion: interface {} is string, not int
// ============================================================
//
// 为什么会这样：
//   Go 提供两种类型断言语法：
//   1. v := x.(T)        —— 断言失败直接 panic（不安全）
//   2. v, ok := x.(T)    —— 断言失败 ok=false（安全！）
//
//   Go 社区强烈建议始终使用 comma-ok 模式。
//
// CONTRAST（与已知语言对比）：
//   - TS:   const x = val as string; // 编译时强制，运行时可能错
//   - Rust: let x = val.downcast_ref::<String>(); // 返回 Option
//   - Go:   val.(string) // 运行时检查，失败 panic（类似 unwrap）

package main

import "fmt"

func main() {
	var data any = "hello"

	// BUG：不安全断言（类型不匹配 → panic）
	// num := data.(int) // ← 取消注释会 panic

	// 安全断言：comma-ok 模式 ✅
	if str, ok := data.(string); ok {
		fmt.Printf("是字符串: %s\n", str)
	} else {
		fmt.Println("不是字符串")
	}

	if num, ok := data.(int); ok {
		fmt.Printf("是整数: %d\n", num)
	} else {
		fmt.Printf("不是整数，data 的类型是 %T\n", data)
	}

	// type switch 更安全
	switch v := data.(type) {
	case string:
		fmt.Printf("type switch: 字符串 \"%s\"\n", v)
	case int:
		fmt.Printf("type switch: 整数 %d\n", v)
	default:
		fmt.Printf("type switch: 其他类型 %T\n", v)
	}
}
