// WHAT: Bug 03 — range 值拷贝陷阱
// ERROR: range 遍历时，第二个返回值是值的拷贝，修改它不会影响原切片
//        这是 Go 中新人常见误区
//
// ============================================================
// 运行结果（取地址）：
// 所有元素地址相同！因为 v 是同一个变量
// ============================================================
//
// 为什么会这样：
//   for i, v := range slice 中的 v 在每次迭代时是切片元素的**拷贝**
//   v 的地址在整个循环中不变（同一个内存位置）
//   修改 v 不会影响原切片元素
//
// CONTRAST（与已知语言对比）：
//   - Rust: for x in &mut vec { *x = ... }   —— 明确是引用
//   - TS:   for (let v of arr) { v = ... }   —— v 是拷贝，不影响原数组
//   - Go:   for _, v := range slice { v = ... } —— 同样不影响原切片！
//
//   关键：Go 的 range 行为与 TS 一致（值拷贝），但 Go 有指针，
//        所以看起来像是引用，实际上不是
//
// 如何修复：
//   方法 1：用索引修改
//     for i := range slice { slice[i] = ... }
//   方法 2：遍历指针切片
//     for _, p := range ptrSlice { *p = ... }

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("原始切片:", nums)

	// BUG 演示：尝试用 range 的值变量修改元素
	for _, v := range nums {
		v = v * 2 // 只修改了拷贝！
	}
	fmt.Println("range 值修改后:", nums) // [1 2 3 4 5] ← 没变！

	// BUG 2：取 range 值变量的地址
	for _, v := range nums {
		fmt.Printf("&v = %p, v = %d\n", &v, v)
		// 所有地址都一样！v 是每次被赋值的同一个变量
	}

	// 修复：用索引修改
	for i := range nums {
		nums[i] = nums[i] * 2
	}
	fmt.Println("索引修改后:", nums) // [2 4 6 8 10] ✅

	// 修复：用指针切片（另一种方式）
	fmt.Println("\n--- 指针切片方式 ---")
	ptrs := []*int{new(int), new(int), new(int)}
	*ptrs[0], *ptrs[1], *ptrs[2] = 10, 20, 30
	fmt.Print("原始: ")
	for _, p := range ptrs {
		fmt.Printf("%d ", *p)
	}
	fmt.Println()
	for _, p := range ptrs {
		*p = *p * 2 // 通过指针修改 ✅
	}
	fmt.Print("修改后: ")
	for _, p := range ptrs {
		fmt.Printf("%d ", *p)
	}
	fmt.Println()
}
