// WHAT: Bug 02 — 零值判断混淆：空字符串 "" vs nil
// ERROR: 误以为 string 类型的零值是 nil，用 nil 判断空字符串
//        或误以为 nil slice 和空 slice 完全一样
//
// ============================================================
// 编译错误信息（关键行）：
// cannot use nil as string value in variable declaration
// ============================================================
//
// 为什么会这样：
//   Go 的 string 是**值类型**（底层是 struct{*byte, int}），不能为 nil。
//   只有 pointer、slice、map、channel、func、interface 可以为 nil。
//   string 的零值是 ""（空字符串），不是 nil。
//
// CONTRAST（与已知语言对比）：
//   - TS:  let s: string | null = null;   // null 合法
//   - Rust: let s: Option<String> = None;  // 用 Option 表达
//   - C:    char* s = NULL;                // 指针可以为 NULL
//   - Go:   var s string                   // s == ""，永远不会 nil
//
//   这种设计的好处：你永远不需要在 Go 中对 string 做 nil 检查，
//   但坏处是：如果你需要区分"空字符串"和"未设置"，就得额外用
//   指针（*string）或特殊值。
//
// 如何修复：
//   判断空字符串：s == ""
//   需要区分"空"和"未设置"：用 *string（指针）或 sql.NullString
//
// ============================================================

package main

import "fmt"

func main() {
	// BUG 1：试图把 nil 赋给 string —— 编译错误！
	// var s string = nil  // ← 取消注释会报：cannot use nil as string value

	var s string // 正确的做法：零值就是 ""
	fmt.Printf("空字符串: %q, 长度为: %d\n", s, len(s))
	// 输出: 空字符串: "", 长度为: 0

	// BUG 2：对 nil map 写操作 —— 运行时会 panic！
	var m map[string]int // nil map
	fmt.Printf("nil map len: %d\n", len(m)) // 读 len 没问题

	// 读 nil map 也 OK，返回零值
	val := m["不存在的键"]
	fmt.Printf("读 nil map 不存在的键: %d\n", val) // 0

	// 但写 nil map 会 panic！
	// m["key"] = 1  // ← 取消注释会 panic: assignment to entry in nil map

	// ============================================================
	// 修复版本
	// ============================================================
	//
	// var m2 = make(map[string]int) // 用 make 初始化
	// m2["key"] = 1                 // 安全
	//
	// var s2 string = ""            // string 零值就是 ""
	// if s2 == "" {
	//     fmt.Println("是空字符串，不是 nil")
	// }
	//

	fmt.Println("✅ 此文件是 bug 演示，编译不会报错但运行取消注释会出问题")
}
