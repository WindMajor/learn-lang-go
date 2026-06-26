// WHAT: Bug 02 — iota 重置陷阱
// ERROR: 每遇到新的 const 块，iota 会重置为 0
//        很多人以为 iota 是全局递增的，实际上每个 const 块独立
//
// ============================================================
// 运行结果：
// Apple = 0
// Banana = 1
// Cat = 0
// Dog = 1
// ============================================================
//
// 为什么会这样：
//   iota 的值在每个 const 声明块开始时重置为 0。
//   每个 const (...) 是一个独立的"枚举组"。
//   Go 设计如此：枚举组之间不应该有依赖关系。
//
// CONTRAST（与已知语言对比）：
//   - Rust: 每个 enum 独立，没有全局递增概念
//   - TS:   每个 enum 独立
//   - Go:   iota 在每个 const 块独立
//
//   这个设计是合理的 —— 想象一下如果 iota 全局递增，
//   修改中间的枚举定义会影响后面所有枚举的值，这是灾难性的。
//
// 如何修复（如果你需要连续的值）：
//   把它们放在同一个 const 块中：
//     const (
//         Apple = iota
//         Banana
//         Cat
//         Dog
//     )
//
// ============================================================

package main

import "fmt"

// BUG: 两个独立的 const 块，iota 都在两个块中从 0 开始
const (
	Apple = iota // 0
	Banana       // 1
)

const (
	Cat = iota // 0 ← 重置了！不是 2
	Dog        // 1
)

func main() {
	fmt.Printf("Apple = %d (期望 0, 实际 %d)\n", Apple, Apple)
	fmt.Printf("Banana = %d (期望 1, 实际 %d)\n", Banana, Banana)
	fmt.Printf("Cat = %d (期望 2, 实际 %d) ← BUG!\n", Cat, Cat)
	fmt.Printf("Dog = %d (期望 3, 实际 %d) ← BUG!\n", Dog, Dog)

	fmt.Println("\niota 在每个 const 块中重置为 0，这是 Go 的设计选择。")
	fmt.Println("如果需要连续值，请放在同一个 const 块中。")
}
