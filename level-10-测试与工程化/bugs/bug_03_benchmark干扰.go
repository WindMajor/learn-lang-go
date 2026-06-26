// WHAT: Bug 03 — Benchmark 中的编译器优化干扰
// ERROR: 在 benchmark 中，编译器可能把"无用"的返回值优化掉
//        导致被测代码实际上没有被执行！
//
// ============================================================
// 运行结果：
// BenchmarkDummy: 0.2 ns/op  ← 不合理的快！
// ============================================================
//
// 为什么会这样：
//   如果 benchmark 代码的结果没有被使用，编译器可能把这个计算
//   完全优化掉。比如 BenchmarkDummy 中 Add 的结果没有被使用，
//   编译器可能直接跳过 Add 调用。
//
// 如何修复：
//   把结果赋值给包级别的变量，防止编译器优化：
//     var result int
//     func BenchmarkFixed(b *testing.B) {
//         var r int
//         for i := 0; i < b.N; i++ {
//             r = Add(10, 20)
//         }
//         result = r  // 赋值给包级变量，阻止优化
//     }

package main

import (
	"testing"
)

func Add(a, b int) int { return a + b }

// BUG: 结果未被使用，可能被编译器优化
func BenchmarkDummy(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Add(10, 20) // ← 结果被丢弃！编译器可能优化掉
	}
}

// 修复: 把结果赋值给包级别变量
var sink int

func BenchmarkFixed(t *testing.B) {
	for i := 0; i < t.N; i++ {
		sink = Add(10, 20) // ← 结果赋值给包级变量，无法优化
	}
}

// 替代方案：使用 testing 包内置的辅助
// BenchmarkFixed(t *testing.B) {
//     for i := 0; i < t.N; i++ {
//         t.SetBytes(8) // 设置每次操作处理的字节数
//         Add(10, 20)
//     }
// }
