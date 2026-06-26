// WHAT: 沙盒文件 — 函数、指针、结构体实验
package main

import "fmt"

func main() {
	fmt.Println("🎮 Level 04 沙盒\n")

	// ======== 多返回值 ========
	fmt.Println("--- 多返回值 ---")
	q, r := div(10, 3)
	fmt.Printf("10/3 = %d 余 %d\n", q, r)

	// ======== 变参 ========
	fmt.Println("--- 变参 ---")
	fmt.Printf("add(1,2,3,4,5) = %d\n", add(1, 2, 3, 4, 5))

	// ======== defer ========
	fmt.Println("--- defer ---")
	defer fmt.Println("第3个")
	defer fmt.Println("第2个")
	fmt.Println("第1个")

	// ======== 指针 ========
	fmt.Println("\n--- 指针 ---")
	x := 42
	px := &x
	fmt.Printf("x=%d, &x=%p, *px=%d\n", x, px, *px)

	// ======== 结构体 ========
	fmt.Println("\n--- 结构体 ---")
	type Point struct{ X, Y int }
	p := Point{10, 20}
	fmt.Printf("Point: %+v\n", p)

	// ======== 你的实验区 ========
	fmt.Println("\n--- 你的实验 ---")

	fmt.Println("✅ 沙盒完成！")
}

func div(a, b int) (int, int) { return a / b, a % b }
func add(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}
