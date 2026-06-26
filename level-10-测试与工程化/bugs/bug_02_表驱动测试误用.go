// WHAT: Bug 02 — 表驱动测试中 t.Run 未正确隔离
// ERROR: 表驱动测试中直接使用 t.Fatal/Error 而不通过 t.Run 子测试隔离
//        导致一个用例失败后，后续用例无法执行
//
// ============================================================
// 运行结果（某个用例失败后）：
// --- FAIL: TestFailFirst
//     bug_test.go:XX: 用例1失败
// (用例2、用例3没有执行记录)
// ============================================================
//
// 为什么会这样：
//   t.Fatal/t.Fatalf 会立即停止当前测试函数。
//   如果不使用 t.Run 隔离，一个用例的 t.Fatal 会终止整个测试函数。
//
// 如何修复：
//   使用 t.Run 让每个用例成为独立的子测试

package main

import "testing"

// BUG: 没有 t.Run，t.Fatal 会终止整个函数
func TestWithoutSubTests(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{1, 2, 99},  // 这个会失败
		{2, 3, 5},   // 这个不会被运行
		{3, 4, 7},   // 这个也不会被运行
	}

	for _, tc := range cases {
		// 使用 t.Fatalf — 如果用例失败，整个测试停止
		if tc.a+tc.b != tc.want {
			t.Fatalf("%d+%d = %d, want %d (后面的用例不会被运行!)", tc.a, tc.b, tc.a+tc.b, tc.want)
		}
	}
}

// 修复: 用 t.Run 隔离
func TestWithSubTests(t *testing.T) {
	cases := []struct {
		name string
		a, b, want int
	}{
		{"第一个（会失败）", 1, 2, 99},
		{"第二个（会成功）", 2, 3, 5},
		{"第三个（会成功）", 3, 4, 7},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// 用 t.Errorf 而非 t.Fatalf，或即使用 Fatalf 也只影响子测试
			if tc.a+tc.b != tc.want {
				t.Errorf("%d+%d = %d, want %d", tc.a, tc.b, tc.a+tc.b, tc.want)
			}
		})
	}
}
