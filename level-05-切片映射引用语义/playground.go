// WHAT: 沙盒文件 — 切片与 map 实验
package main

import "fmt"

func main() {
	fmt.Println("🎮 Level 05 沙盒\n")

	// ======== 切片子切片实验 ========
	fmt.Println("--- 子切片共享 ---")
	a := []int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Printf("a=%v, s=%v\n", a, s)
	s[0] = 999
	fmt.Printf("修改 s[0]=999: a=%v\n\n", a)

	// ======== append 实验 ========
	fmt.Println("--- append ---")
	// 试试：append 什么时候返回新底层数组？
	b := make([]int, 2, 3)
	b[0], b[1] = 1, 2
	fmt.Printf("b: %v, len=%d, cap=%d, ptr=%p\n", b, len(b), cap(b), b)
	b = append(b, 3)
	fmt.Printf("append 3: %v, len=%d, cap=%d, ptr=%p\n", b, len(b), cap(b), b)
	b = append(b, 4)
	fmt.Printf("append 4: %v, len=%d, cap=%d, ptr=%p ← 新地址\n\n", b, len(b), cap(b), b)

	// ======== map 实验 ========
	fmt.Println("--- map ---")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	// 试试删除不存在的 key
	delete(m, "z")
	fmt.Printf("delete 不存在的 key 也没事: %v\n", m)

	// ======== 你的实验区 ========
	fmt.Println("\n--- 你的实验 ---")

	fmt.Println("✅ 沙盒完成！")
}
