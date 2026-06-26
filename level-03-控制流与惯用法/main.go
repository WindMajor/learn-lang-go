// WHAT: Level 03 主代码 — Go 控制流与惯用法
// WHY: 展示 Go 的 if/for/switch/range 用法和设计选择
// CONTRAST: 对比 Rust 的 match/loop、C 的 while、TS 的 for-of

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("========== Go 控制流与惯用法 ==========\n")

	// ============================================================
	// 第一部分：if — 支持短声明
	// ============================================================
	fmt.Println("--- 1. if 与短声明 ---")

	// WHAT: if 可以在条件前执行一个短声明，变量作用域限定在 if-else 块内
	// WHY: 这是 Go 最惯用的模式 —— 把"获取"和"判断"写在一行，减少变量泄露
	// CONTRAST:
	//   - Rust: if let Some(x) = option { }  —— 类似但用于模式匹配
	//   - TS:   无等价物，通常分两行写
	//   - C:    无等价物
	if data := fetchData(); data > 0 {
		fmt.Printf("  ✅ 获取到数据: %d (data 仅在此可见)\n", data)
	} else if data == 0 {
		fmt.Printf("  ⚠️  数据为空: %d\n", data)
	} else {
		fmt.Printf("  ❌ 错误: %d\n", data)
	}
	// data 在这里不可见（作用域限于 if-else 块）

	// ============================================================
	// 第二部分：for — Go 唯一的循环
	// ============================================================
	fmt.Println("\n--- 2. for 的四种形态 ---")

	// 形态 1：标准 C 风格 for
	fmt.Print("  形态 1 (经典 for): ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 形态 2：while 风格（Go 没有 while 关键字）
	// WHY: Go 的设计者认为一种循环语法足够，减少心智负担
	//      但代价是：读代码时不如"while n > 0"直观
	// CONTRAST:
	//   - Rust: while n > 0 { }
	//   - C:    while (n > 0) { }
	//   - Go:   for n > 0 { }  ← 看起来还是 for，语义上是 while
	fmt.Print("  形态 2 (while 风格): ")
	n := 3
	for n > 0 {
		fmt.Printf("%d ", n)
		n--
	}
	fmt.Println()

	// 形态 3：无限循环
	// fmt.Print("  形态 3 (无限循环): for { }  // 取消注释看效果")

	// 形态 4：range 遍历
	// WHAT: range 返回 (索引, 值) 或 (键, 值)
	// WHY: range 是 Go 中遍历任意集合的统一方式
	fmt.Print("  形态 4 (range slice): ")
	items := []string{"Go", "Rust", "TypeScript", "Python"}
	for i, v := range items {
		fmt.Printf("[%d:%s] ", i, v)
	}
	fmt.Println()

	// ============================================================
	// 第三部分：range 的高级用法
	// ============================================================
	fmt.Println("\n--- 3. range 高级用法 ---")

	// range string — 按 rune 遍历（不是按 byte！）
	fmt.Print("  range \"你好Go\": ")
	for i, r := range "你好Go" {
		fmt.Printf("[%d:%c] ", i, r)
	}
	fmt.Println("(注意索引不连续，因为是按 rune 对齐)")

	// range map — 遍历顺序随机！
	// WHAT: Go 故意让 map 遍历顺序随机，防止程序员依赖遍历顺序
	// WHY: map 底层是 hash，迭代顺序不确定；随机化是刻意为之
	// CONTRAST:
	//   - Rust: HashMap 遍历顺序也不确定
	//   - Python: dict 从 3.7 开始保证插入顺序
	//   - TS:   Map 保证插入顺序
	//   - Go:   map 遍历顺序随机！
	fmt.Println("  range map (顺序随机！):")
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for k, v := range m {
		fmt.Printf("    %s:%d", k, v)
	}
	fmt.Println()

	// range 时可以丢弃不需要的值
	fmt.Print("  range 只取值: ")
	for _, v := range items {
		fmt.Printf("%s ", v)
	}
	fmt.Println()

	fmt.Print("  range 只取索引: ")
	for i := range items {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// ============================================================
	// 第四部分：switch — 自动 break
	// ============================================================
	fmt.Println("\n--- 4. switch — 自动 break ---")

	// WHAT: Go 的 switch 默认不穿透（自动 break），不需要写 break
	// WHY: C 的 switch 默认穿透是著名的 Bug 来源，Go 修正了这个设计失误
	//      如果你需要穿透，用 fallthrough 关键字（很少用）
	// CONTRAST:
	//   - C:    switch 默认穿透，需要 break 阻止（命名的 Bug 来源）
	//   - Rust: match 默认可穷尽，不穿透
	//   - TS:   switch 默认穿透，需要 break
	//   - Go:   switch 默认不穿透！

	// 模式 1：表达式 switch
	os := "linux"
	switch os {
	case "darwin":
		fmt.Println("  🍎 macOS")
	case "linux":
		fmt.Println("  🐧 Linux")
	case "windows":
		fmt.Println("  🪟 Windows")
	default:
		fmt.Println("  ❓ 未知系统")
	}

	// 模式 2：无表达式 switch（替代 if-else 链）
	// WHY: 这种写法比 if-else 链更清晰，尤其适合多条件判断
	score := 85
	switch {
	case score >= 90:
		fmt.Println("  等级: A")
	case score >= 80:
		fmt.Println("  等级: B")
	case score >= 70:
		fmt.Println("  等级: C")
	case score >= 60:
		fmt.Println("  等级: D")
	default:
		fmt.Println("  等级: F")
	}

	// 模式 3：switch 支持短声明（和 if 一样）
	switch lang := "Go"; lang {
	case "Go":
		fmt.Println("  ✅ 正在学 Go")
	case "Rust":
		fmt.Println("  你已经会 Rust 了")
	}

	// 模式 4：多值匹配
	switch day := 3; day {
	case 1, 2, 3, 4, 5:
		fmt.Println("  工作日")
	case 6, 7:
		fmt.Println("  周末")
	}

	// fallthrough：强制穿透到下一个 case（很少使用）
	fmt.Print("  fallthrough 演示: ")
	switch n := 1; n {
	case 1:
		fmt.Print("case 1 → ")
		fallthrough
	case 2:
		fmt.Print("case 2 → ")
		fallthrough
	case 3:
		fmt.Println("case 3")
	}

	// ============================================================
	// 第五部分：Go 没有三目运算符
	// ============================================================
	fmt.Println("\n--- 5. 没有三目运算符 ---")

	// WHAT: Go 没有 condition ? a : b 语法
	// WHY: Go 的设计者认为三目运算符让代码难读，if-else 更清晰
	//      但也有人说这是 Go 的固执 —— 你用惯了就会发现 if 也不差
	// CONTRAST:
	//   - C/TS/Python: 有三目 x = a > b ? a : b
	//   - Rust:        没有三目，但 if 是表达式: let x = if a > b { a } else { b };
	//   - Go:          if 是语句，不能嵌入表达式，必须老老实实写 if-else

	a, b := 10, 20
	var max int
	if a > b {
		max = a
	} else {
		max = b
	}
	fmt.Printf("  max(%d, %d) = %d (没有 a > b ? a : b)\n", a, b, max)

	// ============================================================
	// 第六部分：break / continue / goto
	// ============================================================
	fmt.Println("\n--- 6. break / continue / goto ---")

	// break 标签 — 跳出多层循环
	// CONTRAST: Rust 用 loop 标签: 'outer: loop { break 'outer; }
	fmt.Println("  break 标签跳出双层循环：")
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("  在 i=%d, j=%d 处跳出\n", i, j)
				break outer
			}
			fmt.Printf("  (%d,%d) ", i, j)
		}
	}
	fmt.Println()

	// continue 标签 — 跳过外层循环的当前迭代
	fmt.Print("  continue 标签: ")
outer2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue outer2 // 跳过外层当前迭代
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()

	// goto — Go 有 goto，但在现代 Go 中极少使用
	// WHY: Go 保留了 goto 用于自动生成的代码和极少数场景
	//      日常编程中用 goto 是 Go 社区的代码异味（code smell）

	// ============================================================
	// 第七部分：defer 预告
	// ============================================================
	fmt.Println("\n--- 7. defer 预告 ---")
	fmt.Println("  注意：defer 的执行顺序和求值时机将在 Level 04 详细学习")
	fmt.Println("  这里先预告一下：defer 是 Go 的资源管理基石")

	fmt.Println("\n✅ Level 03 完成！")
}

// fetchData 模拟获取数据（用于演示 if 短声明）
func fetchData() int {
	return 42
}

// 以下代码展示了 Go 的 idomatic 错误处理模式（Level 08 详讲）
// 这里先体会结构
func processFile(filename string) error {
	// 模拟：处理文件
	if strings.Contains(filename, "..") {
		return fmt.Errorf("非法文件名: %s", filename)
	}
	return nil
}
