// WHAT: Bug 02 — 命名返回值遮蔽陷阱
// ERROR: 在函数体内用 := 声明了和命名返回值同名的变量
//        这会导致 return 时返回的是命名返回值的零值，而非你赋的值
//
// ============================================================
// 运行结果：
// 返回: 0  ← 不是期望的 10！
// ============================================================
//
// 为什么会这样：
//   命名返回值就是函数体的变量，裸 return 返回这些变量的当前值。
//   如果你在函数体内用 := 声明了同名变量，你就遮蔽（shadow）了
//   命名返回值，裸 return 返回的是外部那个（值是零值）。
//
// CONTRAST（与已知语言对比）：
//   - Rust: 没有命名返回值，不会遇到这个问题
//   - Go:   命名返回值和 := 的作用域规则组合产生了这个陷阱
//
// 如何修复：
//   1. 避免使用裸 return，总是显式 return（Go 社区的最佳实践）
//   2. 如果用命名返回值，内部用 =（不是 :=）来赋值

package main

import "fmt"

// BUG: 命名返回值被遮蔽
func buggyFunction() (result int) {
	// result 是命名返回值，初始为零值 0

	if true {
		result := 10 // ← BUG! 创建了新变量，遮蔽了命名返回值！
		_ = result    // 这个 result 是新的局部变量
	}

	return // 裸 return → 返回的是外层的 result（零值 0）
}

// 修复版本：不要用 := 遮蔽
func fixedFunction() (result int) {
	if true {
		result = 10 // ← 用 = 而非 :=，修改的就是命名返回值
	}
	return result // 或者用裸 return（但显式 return 更好）
}

// 另一个修复：避免命名返回值
func betterFunction() int {
	result := 0 // 局部变量
	if true {
		result = 10
	}
	return result
}

func main() {
	fmt.Printf("buggyFunction() = %d  ← BUG! 期望 10\n", buggyFunction())
	fmt.Printf("fixedFunction() = %d\n", fixedFunction())
	fmt.Printf("betterFunction() = %d\n", betterFunction())
}
