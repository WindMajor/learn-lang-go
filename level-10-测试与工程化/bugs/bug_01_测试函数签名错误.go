// WHAT: Bug 01 — 测试函数签名错误
// ERROR: 测试函数必须以 Test 开头，且参数必须是 *testing.T
//        如果签名不对，go test 不会识别为测试
//
// ============================================================
// 运行 go test：
// testing: warning: no tests to run
// ============================================================
//
// 为什么会这样：
//   Go test 通过函数的命名约定来识别测试：
//   1. 文件名必须是 *_test.go
//   2. 函数名必须以 Test 开头（区分大小写）
//   3. 参数必须是 *testing.T
//   4. 没有返回值
//
// 如何修复：
//   func TestAdd(t *testing.T) { ... }     ← 正确
//   func testAdd(t *testing.T) { ... }     ← 错误！首字母小写
//   func TestAdd() { ... }                 ← 错误！缺少 T 参数

package main

import "testing"

// BUG: 函数名小写 — go test 不会识别
func testWrongName(t *testing.T) {
	// 这个测试永远不会被执行！
	t.Error("不会被执行")
}

// BUG: 缺少 testing.T 参数
// func TestMissingT() { ... }  ← go tool 不认为是测试

// 正确写法
func TestCorrectSignature(t *testing.T) {
	// 这个测试会被执行
	if 1+1 != 2 {
		t.Error("数学崩溃了")
	}
}
