// WHAT: Level 05 主代码 — 切片（slice）、映射（map）与引用语义
// WHY: 理解 Go 数据结构的底层机制，是写出正确 Go 代码的关键
// CONTRAST: Rust 的 Vec 有所有权而 slice 没有，C 的裸指针 vs Go 的安全切片

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("========== 切片、映射与引用语义 ==========\n")

	// ============================================================
	// 第一部分：数组（Array）—— 固定大小，值类型
	// ============================================================
	fmt.Println("--- 1. 数组 vs 切片 ---")

	// WHAT: Go 数组是定长的，长度是类型的一部分
	//       [3]int 和 [4]int 是不同的类型！
	// WHY: 数组在 Go 中是值类型，赋值会拷贝整个数组（不像 C 退化为指针）
	// CONTRAST:
	//   - C:    int arr[3]; 数组名退化为指针，sizeof(arr) == sizeof(&arr[0])
	//   - Rust: let arr: [i32; 3] = [1, 2, 3]; 数组大小也是类型一部分
	//   - Go:   arr := [3]int{1, 2, 3} —— 数组大小是类型一部分，传参会拷贝！

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("  数组: %v, len=%d\n", arr, len(arr))
	fmt.Printf("  数组地址: %p\n", &arr)
	fmt.Printf("  数组元素地址: %p, %p, %p\n", &arr[0], &arr[1], &arr[2])
	// 连续内存！和 C 一样（但自动管理）

	arr2 := arr // 值拷贝！arr2 是 arr 的完整副本
	arr2[0] = 999
	fmt.Printf("  arr  %v (修改 arr2 不影响 arr)\n", arr)
	fmt.Printf("  arr2 %v\n\n", arr2)

	// ============================================================
	// 第二部分：切片（Slice）—— 动态视图
	// ============================================================
	fmt.Println("--- 2. Slice 底层结构 ---")

	// WHAT: slice 底层是一个 struct{ptr, len, cap}
	//       ptr 指向底层数组，多个 slice 可以共享同一个底层数组！
	// WHY: 提供动态数组能力，同时保持零分配（子切片不拷贝数据）
	// CONTRAST:
	//   - Rust: Vec<T> 拥有数据，&[T] 是借用（borrow）
	//   - Go:   slice 类似 &[T]，但 Go 没有借用检查器！
	//           这意味着多个 slice 共享底层数组时，修改一个会影响另一个

	// 创建 slice 的方式
	s1 := []int{1, 2, 3, 4, 5}      // 字面量（底层自动创建数组）
	s2 := make([]int, 3, 5)           // make: len=3, cap=5（底层数组有 5 个位置）
	var s3 []int                       // nil slice（ptr=nil, len=0, cap=0）

	fmt.Printf("  s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("  s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("  s3: %v, len=%d, cap=%d, isNil=%t\n\n", s3, len(s3), cap(s3), s3 == nil)

	// ============================================================
	// 第三部分：子切片与底层数组共享（关键！）
	// ============================================================
	fmt.Println("--- 3. 子切片与底层数组共享 ---")

	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sub := original[2:5] // 索引 2~4，共享底层数组！
	fmt.Printf("  original: %v\n", original)
	fmt.Printf("  sub[2:5]: %v\n\n", sub)

	// 修改 sub，original 也变了！
	sub[0] = 999
	fmt.Printf("  修改 sub[0]=999 后:\n")
	fmt.Printf("  original: %v ← 也变了！\n", original)
	fmt.Printf("  sub:      %v\n\n", sub)

	// 这就像 Rust 中：let sub = &mut original[2..5]; sub[0] = 999;
	// 但 Go 没有借用检查，你很容易无意中修改了"别人的"数据

	// ============================================================
	// 第四部分：append 与扩容
	// ============================================================
	fmt.Println("--- 4. append 与扩容 ---")

	// WHAT: append 可能返回新切片（如果 cap 不够）
	//       扩容后新切片有独立的底层数组
	// WHY: 这是 Go slice 自动内存管理的核心机制
	// 规则：cap < 1024 时，扩容为 2×cap；cap ≥1024 时，扩容 1.25×cap

	s := make([]int, 0, 2) // len=0, cap=2
	fmt.Printf("  初始: len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

	s = append(s, 1)
	fmt.Printf("  append(1):  len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

	s = append(s, 2)
	fmt.Printf("  append(2):  len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

	s = append(s, 3) // cap 不够，扩容！
	fmt.Printf("  append(3):  len=%d, cap=%d, ptr=%p ← 新地址！\n\n", len(s), cap(s), s)

	// 关键认识：append 总是返回一个值，你必须接收它！
	// append(s, x) // ← 这个操作如果没有 s = ...，修改白做了！

	// ============================================================
	// 第五部分：copy —— 真正的深拷贝
	// ============================================================
	fmt.Println("--- 5. copy 深拷贝 ---")

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copied := copy(dst, src) // 返回拷贝的元素个数
	fmt.Printf("  copy 了 %d 个元素，dst: %v\n", copied, dst)
	fmt.Printf("  dst[0]地址=%p, src[0]地址=%p ← 不同！\n\n", &dst[0], &src[0])

	// ============================================================
	// 第六部分：nil slice vs 空 slice
	// ============================================================
	fmt.Println("--- 6. nil slice vs 空 slice ---")

	var nilSlice []int         // nil slice
	emptySlice := []int{}      // 空 slice（非 nil）
	makeSlice := make([]int, 0) // 空 slice（非 nil）

	fmt.Printf("  nil slice:    %v, len=%d, cap=%d, isNil=%t\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("  empty slice:  %v, len=%d, cap=%d, isNil=%t\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("  make slice:   %v, len=%d, cap=%d, isNil=%t\n\n", makeSlice, len(makeSlice), cap(makeSlice), makeSlice == nil)

	// 它们功能性相同（都可以 range、append），但 JSON 序列化不同：
	// nil slice → JSON null，empty slice → JSON []
	// 所以：如果你要返回空列表，用 make([]T, 0) 或 []T{} 而不是 nil

	// ============================================================
	// 第七部分：Map（映射）
	// ============================================================
	fmt.Println("--- 7. Map ---")

	// WHAT: map 是引用类型，底层是 hash 表
	//       map 不是并发安全的！并发读写会 panic（甚至直接崩溃）
	// WHY: Go 选择了让 map 简单快速，并发安全交给开发者（sync.Mutex / sync.Map）
	// CONTRAST:
	//   - Rust: HashMap 也不并发安全，但编译器阻止你共享（所有权+Send/Sync）
	//   - TS:   Map 是引用，也无内置并发安全
	//   - Go:   map 运行时检测并发写 → fatal error: concurrent map writes

	// 创建 map
	m1 := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	m2 := make(map[string]int) // 空 map（可写入）
	var m3 map[string]int       // nil map（写会 panic！）

	fmt.Printf("  m1: %v\n", m1)
	fmt.Printf("  m2: %v, isNil=%t\n", m2, m2 == nil)
	fmt.Printf("  m3: %v, isNil=%t\n\n", m3, m3 == nil)

	// Map 基本操作
	age, ok := m1["Alice"] // comma-ok 惯用法：判断 key 是否存在
	fmt.Printf("  Alice 存在? %t, 年龄: %d\n", ok, age)

	age2, ok2 := m1["Charlie"]
	fmt.Printf("  Charlie 存在? %t, 值: %d (不存在的 key 返回零值)\n\n", ok2, age2)

	// Map 删除
	delete(m1, "Bob")
	fmt.Printf("  删除 Bob 后: %v\n\n", m1)

	// Map 遍历（顺序随机！）
	fmt.Println("  Map 遍历（顺序随机）：")
	for k, v := range m1 {
		fmt.Printf("    %s: %d\n", k, v)
	}
	fmt.Println()

	// ============================================================
	// 第八部分：Slice 常用操作
	// ============================================================
	fmt.Println("--- 8. Slice 常用操作 ---")

	// 删除索引 i 的元素（保持顺序）
	nums := []int{10, 20, 30, 40, 50}
	i := 2 // 删除索引 2（值为 30）
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Printf("  删除索引 %d: %v\n", i, nums)

	// 在索引 i 处插入
	nums = append(nums[:1], append([]int{15}, nums[1:]...)...)
	fmt.Printf("  在索引 1 插入 15: %v\n", nums)

	// 过滤（Go 没有内置 filter）
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filtered := filterSlice(nums, func(x int) bool { return x%2 == 0 })
	fmt.Printf("  过滤偶数: %v\n", filtered)

	// 排序
	unsorted := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sort.Ints(unsorted)
	fmt.Printf("  排序: %v\n\n", unsorted)

	fmt.Println("✅ Level 05 完成！")
}

// filterSlice 切片过滤（Go 没有泛型 filter，需要手动实现或 Go 1.18+ 用泛型）
func filterSlice(s []int, fn func(int) bool) []int {
	var result []int
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
