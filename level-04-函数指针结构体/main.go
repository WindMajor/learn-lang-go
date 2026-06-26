// WHAT: Level 04 主代码 — 函数、指针与结构体
// WHY: Go 的函数设计、defer 资源管理、指针哲学、struct 为中心的数据模型
// CONTRAST: 对比 Rust 的 Result、C 的指针运算、TS 的 class

package main

import (
	"fmt"
	"os"
)

// Person 结构体 —— 演示 struct 的用法
// WHAT: 导出类型，首字母大写（其他包可见）
type Person struct {
	Name string
	Age  int
	Addr string // 注意：Go 没有逗号分隔符（用换行分隔）
}

func main() {
	fmt.Println("========== Go 函数、指针与结构体 ==========\n")

	// ============================================================
	// 第一部分：函数声明 — 多返回值
	// ============================================================
	fmt.Println("--- 1. 多返回值 ---")
	// CONTRAST:
	//   - Rust: 用 Result<T, E> 或 Option<T> 表达可失败，多值用元组
	//   - TS:   用数组/对象解构模拟多返回值
	//   - Go:   原生多返回值，这是 Go 的杀手特性之一
	q, r := divide(10, 3)
	fmt.Printf("  divide(10, 3) = 商: %d, 余: %d\n", q, r)

	// 多返回值 + error（Go 惯用模式）
	_, err := divideCheck(10, 0)
	if err != nil {
		fmt.Printf("  divideCheck(10, 0) → 错误: %v\n", err)
	}

	result, err := divideCheck(10, 3)
	if err != nil {
		fmt.Printf("  错误: %v\n", err)
	} else {
		fmt.Printf("  divideCheck(10, 3) = %d\n", result)
	}

	// ============================================================
	// 第二部分：命名返回值
	// ============================================================
	fmt.Println("\n--- 2. 命名返回值 ---")
	// WHAT: 返回值可以命名，函数体就像有预声明的变量
	// WHY: 主要优势：1) 文档作用（一眼看出返回什么） 2) 配合 defer 修改返回值
	x, y := split(17)
	fmt.Printf("  split(17) = %d, %d\n", x, y)

	// 命名返回值 + defer 的组合（重要模式！）
	fmt.Printf("  namedReturnWithDefer() = %d\n", namedReturnWithDefer())
	// 注意：defer 修改了返回值！这是 Go 的强大模式

	// ============================================================
	// 第三部分：变参函数
	// ============================================================
	fmt.Println("\n--- 3. 变参函数 ---")
	// WHAT: ...Type 表示可变数量参数，函数内作为切片使用
	fmt.Printf("  sum(1, 2, 3) = %d\n", sum(1, 2, 3))
	fmt.Printf("  sum(10, 20, 30, 40, 50) = %d\n", sum(10, 20, 30, 40, 50))

	// 展开切片传给变参函数
	numbers := []int{1, 2, 3, 4}
	fmt.Printf("  sum(numbers...) = %d\n", sum(numbers...))

	// ============================================================
	// 第四部分：函数是一等公民
	// ============================================================
	fmt.Println("\n--- 4. 函数是一等公民 ---")
	// CONTRAST:
	//   - TS: 函数是一等公民，非常自然 (lambda/arrow fn)
	//   - Rust: 闭包是 Fn/FnMut/FnOnce trait，一等公民但有所有权约束
	//   - Go: 函数是一等公民，但 Go 的习惯用法更偏"方法"（method）
	//         闭包也支持，但 Go 社区不滥用闭包

	// 函数赋值给变量
	double := func(x int) int { return x * 2 }
	fmt.Printf("  double(21) = %d\n", double(21))

	// 高阶函数
	numbers2 := []int{1, 2, 3, 4, 5}
	doubled := mapInts(numbers2, double)
	fmt.Printf("  mapInts([1 2 3 4 5], double) = %v\n", doubled)

	// ============================================================
	// 第五部分：defer — Go 的资源管理基石
	// ============================================================
	fmt.Println("\n--- 5. defer — 资源管理基石 ---")
	// WHAT: defer 延迟执行，用于资源释放（类似 RAII/Drop）
	// WHY: Go 没有析构函数，defer 是释放资源的唯一保证手段
	//      配合 defer 和 if err != nil，构成 Go 的资源管理范式
	// CONTRAST:
	//   - Rust: Drop trait 自动释放（RAII），更安全
	//   - C:   手动释放，容易遗漏
	//   - TS:   try-finally（或 using 声明）
	//   - Go:   defer 是显式的"finally"，必须手写

	// 演示 1：defer 的 LIFO 执行顺序
	fmt.Println("  defer LIFO 演示：")
	defer fmt.Println("    ③ defer 3 (最先声明，最后执行)")
	defer fmt.Println("    ② defer 2")
	defer fmt.Println("    ① defer 1 (最后声明，最先执行)")

	// 演示 2：defer 参数在声明时求值
	fmt.Println("\n  defer 参数求值时机演示：")
	i := 1
	defer fmt.Printf("    defer 时 i = %d (声明时求值！)\n", i)
	i = 100
	fmt.Printf("    修改后 i = %d\n", i)
	// 输出：修改后 i = 100  →  defer 时 i = 1  ← 注意！

	// 演示 3：defer + 命名返回值（重要模式！）
	// 见 namedReturnWithDefer 函数

	// 演示 4：defer 在循环中（性能陷阱，见 bugs/）

	// ============================================================
	// 第六部分：指针 — 有但无运算
	// ============================================================
	fmt.Println("\n--- 6. 指针 — 有但无运算 ---")
	// WHAT: Go 有指针，但不能做算术运算（p++、p+1 都不允许）
	// WHY: 保留指针的零成本抽象价值，同时杜绝 C 的内存安全隐患
	// CONTRAST:
	//   - C:    int* p = &x; p++; *(p+2) // 自由运算，容易越界
	//   - Rust: let p: *mut i32 = &mut x; // raw pointer 可运算但仅 unsafe 中
	//          let r: &mut i32 = &mut x;  // 引用不能运算
	//   - Go:   var p *int = &x; *p = 10  // 只能取地址和解引用，不能运算！

	a := 42
	p := &a     // p 是指向 a 的指针
	fmt.Printf("  a = %d, &a = %p, p = %p\n", a, &a, p)
	fmt.Printf("  通过指针读: *p = %d\n", *p)

	*p = 100    // 通过指针修改
	fmt.Printf("  通过指针写后: a = %d\n", a)

	// Go 指针 vs C 指针 的关键差异
	//   1. Go: 不能 p++（编译错误）
	//   2. Go: 不能 p + n（编译错误）
	//   3. Go: 数组名不是指针（和 C 的根本区别）
	//   4. Go: 指针零值是 nil（不是随机地址！）

	// nil 指针解引用会 panic（不是 segfault，是 panic）
	var nilPtr *int
	fmt.Printf("  nil 指针: %v\n", nilPtr)
	// *nilPtr = 10 // ← panic: runtime error: invalid memory address

	// ============================================================
	// 第七部分：struct — Go 的数据中心
	// ============================================================
	fmt.Println("\n--- 7. struct — Go 的数据中心 ---")
	// WHAT: Go 没有 class，struct 是唯一的数据组合方式
	//       struct 可以组合（嵌套）、可以有方法（方法定义在外边）
	// WHY: Go 选择组合而非继承，struct + interface 比 class 更灵活
	// CONTRAST:
	//   - Rust: struct 是数据，impl 是方法（和 Go 很像）
	//   - TS:   class 有数据+方法+继承
	//   - C:    struct 纯数据
	//   - Go:   struct 是数据，方法是绑定到类型上的函数（写在 struct 外面！）

	// 定义和创建结构体（Person 类型在包级别定义，此处直接使用）

	// 方式 1：按顺序初始化（不推荐，可读性差）
	p1 := Person{"张三", 30, "北京"}

	// 方式 2：按字段名初始化（推荐）
	p2 := Person{
		Name: "李四",
		Age:  25,
		Addr: "上海",
	}

	// 方式 3：零值 + 逐字段赋值
	var p3 Person
	p3.Name = "王五"
	p3.Age = 35

	fmt.Printf("  p1: %+v\n", p1)
	fmt.Printf("  p2: %+v\n", p2)
	fmt.Printf("  p3: %+v\n", p3)

	// 结构体嵌套（组合，不是继承！）
	type Employee struct {
		Person    // 匿名字段（嵌入）→ 字段提升
		Company string
		Salary  float64
	}

	emp := Employee{
		Person:  Person{Name: "赵六", Age: 28, Addr: "深圳"},
		Company: "Tencent",
		Salary:  50000,
	}
	// 字段提升：可以直接访问嵌入类型的字段
	fmt.Printf("  emp.Name = %s, emp.Company = %s\n", emp.Name, emp.Company)

	// ============================================================
	// 第八部分：值传递 vs 指针传递
	// ============================================================
	fmt.Println("\n--- 8. 值传递 vs 指针传递 ---")
	// WHAT: Go 所有函数参数都是值传递（pass by value）
	//       指针也是按值传递的（传递的是地址的拷贝）
	// WHY:  简单一致——没有"引用传递"这个概念
	// CONTRAST:
	//   - C:   值传递 + 指针（手动管理）
	//   - Rust: 所有权传递（let 转移所有权），引用传递（&）
	//   - Go:   全是值传递，指针也是地址值的拷贝！

	person := Person{Name: "原始", Age: 0}
	fmt.Printf("  修改前: %+v\n", person)

	modifyValue(person) // 传值 → 修改副本
	fmt.Printf("  值传递后: %+v (未变！因为传的是副本)\n", person)

	modifyPointer(&person) // 传指针 → 修改原值
	fmt.Printf("  指针传递后: %+v (改变了！因为通过指针修改)\n", person)

	// 但注意：map/slice 是特殊情况（它们是引用类型的值）
	// map 本身是指针包装器，传 map 就是传指针的拷贝（底层数据共享）
	// 详见 Level 05

	// ============================================================
	// 第九部分：文件操作 with defer
	// ============================================================
	fmt.Println("\n--- 9. defer 实用模式：文件操作 ---")

	// WHAT: Go 标准模式：open → check error → defer close
	// WHY: 确保无论函数如何退出，资源都会被释放
	filename := "/tmp/go-learn-test.txt"
	writeAndReadDemo(filename)

	fmt.Println("\n✅ Level 04 完成！")
}

// divide 多返回值示例（不带 error）
func divide(a, b int) (int, int) {
	if b == 0 {
		return 0, 0 // 多返回值
	}
	return a / b, a % b
}

// divideCheck 多返回值 + error（Go 惯用模式）
// CONTRAST: Rust 用 Result<i32, Error>，Go 用手动检查 error
func divideCheck(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为 0") // Go 的 error 是普通接口值
	}
	return a / b, nil
}

// split 命名返回值 —— 返回值名就是函数体内的变量
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 裸 return → 返回 x, y
	// CONTRAST:
	//   - Rust: 没有裸 return，必须显式 (x, y)
	//   - Go:   naked return 简洁但有争议，长函数中降低可读性
}

// namedReturnWithDefer defer 修改命名返回值（关键模式）
func namedReturnWithDefer() (result int) {
	defer func() {
		result++ // defer 可以修改命名返回值！
		fmt.Printf("    defer 修改后 result = %d\n", result)
	}()
	return 10 // result 变成 10，然后 defer 执行，result++，最终返回 11
}

// sum 变参函数
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// mapInts 高阶函数 —— 函数作为参数
func mapInts(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

// modifyValue 值传递 —— 不修改原值
func modifyValue(p Person) {
	p.Name = "被修改了"
	p.Age = 999
}
// 等价于 Rust 的 fn modify(p: Person) —— 获取所有权，不影响原变量

// modifyPointer 指针传递 —— 修改原值
func modifyPointer(p *Person) {
	p.Name = "被修改了（通过指针）"
	p.Age = 999
}
// 等价于 Rust 的 fn modify(p: &mut Person)

// writeAndReadDemo 文件写入和读取演示 defer 模式
func writeAndReadDemo(filename string) {
	// WHAT: Go 标准模式：
	//   1. 打开资源
	//   2. 立即检查错误
	//   3. defer 释放资源
	// WHY: 简洁且安全——资源释放紧挨着创建，不会忘记
	// CONTRAST:
	//   - Rust: File::open 返回 Result，Drop 自动关闭
	//   - Go:   手动 defer close —— 更显式但容易忘

	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("  创建文件失败: %v\n", err)
		return
	}
	defer f.Close() // ← 资源释放紧挨创建，Go 风格

	_, err = f.WriteString("Hello, Go defer!")
	if err != nil {
		fmt.Printf("  写入失败: %v\n", err)
		return
	}
	fmt.Printf("  ✅ 已写入: %s\n", filename)
}
