// WHAT: 沙盒文件 —— 自由实验变量、常量、iota
// 用法：go run playground.go

package main

import "fmt"

func main() {
	fmt.Println("🎮 Level 02 沙盒\n")

	// ======== 实验 1：所有声明方式 ========
	fmt.Println("--- 实验 1: 声明方式 ---")
	var a int
	var b int = 10
	var c = 20
	d := 30
	var e, f int = 40, 50
	g, h := 60, "hello"

	fmt.Printf("a=%d, b=%d, c=%d, d=%d\n", a, b, c, d)
	fmt.Printf("e=%d, f=%d, g=%d, h=%s\n\n", e, f, g, h)

	// ======== 实验 2：iota 创意用法 ========
	fmt.Println("--- 实验 2: iota 创意 ---")
	const (
		Monday = iota + 1 // 从 1 开始
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Printf("星期: %d, %d, %d, %d, %d, %d, %d\n\n",
		Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	// ======== 实验 3：变量遮蔽实验 ========
	fmt.Println("--- 实验 3: 变量遮蔽 ---")
	count := 1
	fmt.Printf("外部 count = %d\n", count)
	{
		count := 100 // 新变量遮蔽
		fmt.Printf("  内部 count = %d\n", count)
	}
	fmt.Printf("外部 count 仍然是 = %d (未被内部修改)\n\n", count)

	// ======== 实验 4：零值结构体 ========
	fmt.Println("--- 实验 4: 结构体零值 ---")
	type Point struct{ X, Y int }
	var p Point
	fmt.Printf("Point 零值: (%d, %d)\n", p.X, p.Y)
	p.X = 10 // 零值结构体可以直接使用！
	fmt.Printf("修改后: (%d, %d)\n\n", p.X, p.Y)

	// ======== 你的实验区 ========
	fmt.Println("--- 你的实验 ---")

	// === START ===

	// === END ===

	fmt.Println("\n✅ 沙盒完成！")
}
