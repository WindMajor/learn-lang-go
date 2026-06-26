// WHAT: 沙盒文件 —— 供学习者随意修改、破坏、实验
// WHY: 把 main.go 的知识点拆开揉碎，随意改不会破坏主线代码
// 用法：go run playground.go
// 建议：每改一行就重新运行，观察输出变化

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("🎮 沙盒模式 —— 随意修改下面代码，go run playground.go 看效果\n")

	// ======== 实验 1：试试不同类型的零值 ========
	fmt.Println("--- 实验 1: 零值 ---")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("int 零值: %d, float64 零值: %f, bool 零值: %t, string 零值: %q\n\n", i, f, b, s)

	// ======== 实验 2：试试 := 推导的类型 ========
	fmt.Println("--- 实验 2: 类型推导 ---")
	v1 := 42          // 猜猜什么类型？
	v2 := 3.14        // 猜猜什么类型？
	v3 := "hello"     // 显然
	v4 := 42 + 0i     // 猜猜？
	fmt.Printf("v1: %T = %v\n", v1, v1)
	fmt.Printf("v2: %T = %v\n", v2, v2)
	fmt.Printf("v3: %T = %v\n", v3, v3)
	fmt.Printf("v4: %T = %v\n\n", v4, v4)

	// ======== 实验 3：类型转换实验 ========
	fmt.Println("--- 实验 3: 类型转换 ---")
	var a int = 100
	var b8 int8 = int8(a)   // 窄化转换，可能溢出！
	var b32 int32 = int32(a) // 宽化转换，安全
	var fu uint = uint(a)    // 有符号→无符号
	fmt.Printf("int: %d, int8: %d, int32: %d, uint: %d\n\n", a, b8, b32, fu)

	// ======== 实验 4：字符串遍历实验 ========
	fmt.Println("--- 实验 4: 中文字符串 ---")
	text := "Go语言学习"
	fmt.Printf("字符串: %s\n", text)
	fmt.Printf("字节数: %d\n", len(text))
	fmt.Printf("字符数: %d\n", utf8.RuneCountInString(text))
	for idx, r := range text {
		fmt.Printf("  位置 %d: %c (U+%04X)\n", idx, r, r)
	}
	fmt.Println()

	// ======== 实验 5：nil 实验 ========
	fmt.Println("--- 实验 5: nil ---")
	var nilSlice []int
	fmt.Printf("nil 切片: %v, len=%d, isNil=%t\n", nilSlice, len(nilSlice), nilSlice == nil)
	// 试试对 nil slice 做 range，会怎样？
	fmt.Print("range nil slice: ")
	for _, v := range nilSlice {
		fmt.Print(v, " ")
	}
	fmt.Println("(空，不 panic)")
	fmt.Println()

	// ======== 实验 6：你的自由实验区 ========
	// 在下面随意写 Go 代码，测试你的猜想
	// 改坏了也不怕，git checkout 即可恢复

	// === 你的实验代码 START ===

	// === 你的实验代码 END ===

	fmt.Println("✅ 沙盒实验完成！")
}
