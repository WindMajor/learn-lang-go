// WHAT: Bug 03 — append 未接收返回值
// ERROR: append 返回新切片（可能改变了底层数组），如果不接收返回值，
//        原始切片不会更新！
//
// ============================================================
// 运行结果：
// 未接收返回值的切片: [1 2]
// ============================================================
//
// 为什么会这样：
//   append 在 cap 够用时会在原底层数组上写，返回值是同一个 slice 结构体
//   （但 len 变了）。在 cap 不够时，会分配新底层数组。
//   无论哪种情况，你都必须用 s = append(s, x) 的模式接收返回值。
//
// CONTRAST（与已知语言对比）：
//   - Rust: v.push(x); // push 修改自身（&mut self）
//   - TS:   arr.push(x); // push 修改自身（）
//   - Python: lst.append(x); // 修改自身
//   - Go:   s = append(s, x) // 必须接收返回值！
//
// 如何修复：
//   总是写：s = append(s, x)  而不是  append(s, x)

package main

import "fmt"

func main() {
	s := make([]int, 2, 4) // len=2, cap=4
	s[0], s[1] = 1, 2
	fmt.Printf("初始: s=%v, len=%d, cap=%d\n", s, len(s), cap(s))

	// BUG：append 不接收返回值
	append(s, 3) // ← 返回值被丢弃！s 不变
	fmt.Printf("append(s,3) 不接收: s=%v, len=%d ← 没变！\n", s, len(s))

	// 正确：接收返回值
	s = append(s, 3)
	fmt.Printf("s = append(s,3): s=%v, len=%d ← 变了！\n", s, len(s))

	s = append(s, 4)
	fmt.Printf("s = append(s,4): s=%v, len=%d, cap=%d\n", s, len(s), cap(s))

	// cap 不够时 append 会分配新底层数组
	s = append(s, 5) // cap 从 4 变成 8（扩容）
	fmt.Printf("s = append(s,5): s=%v, len=%d, cap=%d ← 扩容了\n", s, len(s), cap(s))
}
