// WHAT: 沙盒文件 — 控制流实验
package main

import "fmt"

func main() {
	fmt.Println("🎮 Level 03 沙盒\n")

	// ======== 实验 1：if 短声明 ========
	fmt.Println("--- if 短声明 ---")
	if x := 100; x > 50 {
		fmt.Printf("x=%d > 50\n", x)
	}

	// ======== 实验 2：for 无限循环（注释掉避免卡住） ========
	fmt.Println("--- for 四种形态 ---")
	// 标准
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	// while 风格
	n := 3
	for n > 0 {
		fmt.Printf("%d ", n)
		n--
	}
	fmt.Println()

	// ======== 实验 3：switch ========
	fmt.Println("--- switch ---")
	word := "hello"
	switch word {
	case "hello":
		fmt.Println("English")
	case "你好":
		fmt.Println("中文")
	default:
		fmt.Println("?")
	}

	// ======== 实验 4：range 实验 ========
	fmt.Println("--- range ---")
	nums := []int{10, 20, 30, 40}
	for i, v := range nums {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	// ======== 你的实验区 ========
	fmt.Println("--- 你的实验 ---")
	// === START ===

	// === END ===

	fmt.Println("\n✅ 沙盒完成！")
}
