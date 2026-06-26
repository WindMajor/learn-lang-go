// WHAT: Bug 01 — 切片共享底层数组陷阱
// ERROR: 子切片和原切片共享底层数组，修改一个会影响另一个
//        这在前端/动态语言开发者看来像是"引用"，实际是底层内存共享
//
// ============================================================
// 运行结果：
// 修改子切片[0] = 999 后，原切片[2]也变成了 999
// ============================================================
//
// 为什么会这样：
//   Go 的 slice 是 struct{ptr, len, cap}，子切片拷贝了这个 struct，
//   但 ptr 指向同一个底层数组。所以子切片修改元素 = 修改底层数组。
//
// CONTRAST（与已知语言对比）：
//   - Rust: let s = &v[1..3]; v[1] = X; // 编译错误！借用规则阻止
//   - TS:   const s = arr.slice(1, 3); s[0] = X; // 不影响 arr（浅拷贝）
//   - Python: s = lst[1:3]; s[0] = X; // 不影响 lst（浅拷贝）
//   - Go:   子切片和原切片共享底层数组！最容易踩坑
//
// 如何修复：
//   需要独立副本时用 copy：
//     newSlice := make([]int, len(sub))
//     copy(newSlice, sub)

package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}

	// 创建子切片
	sub := original[1:3] // 包含 original[1] 和 original[2]

	fmt.Printf("原始切片: %v\n", original)
	fmt.Printf("子切片:   %v\n\n", sub)

	// BUG：你以为只修改了 sub，但 original 也变了！
	sub[0] = 999
	fmt.Printf("修改 sub[0] = 999 后:\n")
	fmt.Printf("原始切片: %v ← 也变了！\n", original)
	fmt.Printf("子切片:   %v\n\n", sub)

	// 为什么会发生：底层数组是这样的
	// original → [1, 2, 3, 4, 5]
	// sub      →    [9, 3]  ← 同一个内存！

	// ============================================================
	// 修复版本：用 copy 创建独立副本
	// ============================================================
	original2 := []int{1, 2, 3, 4, 5}
	sub2 := make([]int, 2)
	copy(sub2, original2[1:3]) // 真正拷贝数据
	sub2[0] = 999

	fmt.Println("修复：使用 copy 创建独立副本")
	fmt.Printf("原始切片: %v ← 不变！\n", original2)
	fmt.Printf("子切片:   %v\n\n", sub2)

	// 额外注意事项：append 可能触发扩容
	// 如果 append 时 cap 不够，Go 分配新底层数组，子切片就不再共享
	fmt.Println("注意：append 触发扩容会创建新底层数组，断开共享")
	s := make([]int, 2, 3) // len=2, cap=3
	s[0], s[1] = 1, 2
	sub3 := s[:] // 共享
	s = append(s, 3)
	fmt.Printf("append 后 s: %v\n", s)
	fmt.Printf("sub3: %v (共享了前 2 个元素)\n", sub3)
	s = append(s, 4) // 触发扩容
	s[0] = 999
	fmt.Printf("扩容后 s: %v\n", s)
	fmt.Printf("sub3: %v ← 不再共享！\n", sub3)
}
