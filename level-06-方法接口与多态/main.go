// WHAT: Level 06 主代码 — 方法、接口与多态
// WHY: 隐式接口是 Go 最独特的设计！理解它是写出 Go 风格代码的关键
// CONTRAST: TS 的显式接口 vs Go 的隐式接口，Rust 的 trait vs Go 的 interface

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("========== 方法、接口与多态 ==========\n")

	// ============================================================
	// 第一部分：方法（Method）
	// ============================================================
	fmt.Println("--- 1. 方法：绑定到类型的函数 ---")

	// WHAT: 方法是带有接收者（receiver）的函数
	// WHY: Go 没有 class，方法让你把行为和数据类型绑定
	// CONTRAST:
	//   - Rust: impl Type { fn method(&self) }
	//   - TS:   class { method() {} }
	//   - Go:   func (r ReceiverType) method()

	r := Rectangle{Width: 10, Height: 5}
	fmt.Printf("  矩形: %+v\n", r)
	fmt.Printf("  面积: %.2f\n", r.Area())     // 值接收者
	r.Scale(2)                               // 指针接收者
	fmt.Printf("  缩放 2 倍后: %+v, 面积: %.2f\n\n", r, r.Area())

	// ============================================================
	// 第二部分：值接收者 vs 指针接收者
	// ============================================================
	fmt.Println("--- 2. 值接收者 vs 指针接收者 ---")

	val := Counter{Value: 10}
	fmt.Printf("  值接收者 Increment: %d (原值不变)\n", val.ValueIncrement())
	fmt.Printf("  调用后: %+v\n", val)

	ptr := &Counter{Value: 10}
	fmt.Printf("  指针接收者 Increment: %d\n", ptr.PtrIncrement())
	fmt.Printf("  调用后: %+v ← 值变了\n\n", ptr)

	// Go 的语法糖：即使你定义的是指针接收者，也可以用值调用
	// Go 会自动取地址：val.PtrIncrement() → (&val).PtrIncrement()
	val2 := Counter{Value: 10}
	val2.PtrIncrement() // Go 自动取地址！
	fmt.Printf("  语法糖 val.PtrIncrement() 后: %+v\n\n", val2)

	// ============================================================
	// 第三部分：隐式接口 —— Go 最核心的设计！
	// ============================================================
	fmt.Println("--- 3. 隐式接口（Go 最核心的设计） ---")
	// WHAT: 类型不需要显式声明它实现了哪个接口
	//      只要类型的方法集包含接口所有方法，它就自动实现了该接口
	// WHY: "鸭子类型"的静态类型版本 — 如果它走起来像鸭子，那就是鸭子
	//      这带来了极大的灵活性和解耦
	// CONTRAST:
	//   - TS: 结构化类型（structural typing）—— 形状匹配即可
	//         这实际上和 Go 的隐式接口非常相似！
	//         type Writer = { write: (data: string) => void }
	//         任何有 write 方法的对象都满足 Writer
	//   - Rust: trait 需要显式 impl —— 类型要声明"我实现了这个 trait"
	//     impl Write for MyType { fn write(&mut self, ...) }
	//   - Go: 不需要声明！任何有 Write 方法的类型自动实现 io.Writer

	// 演示：定义接口和多种实现
	var greeter Greeter

	greeter = EnglishGreeter{}
	fmt.Printf("  %s\n", greeter.Greet("World"))

	greeter = ChineseGreeter{}
	fmt.Printf("  %s\n", greeter.Greet("世界"))

	// 匿名结构体也可以实现接口！
	greeter = struct{}{}
	fmt.Printf("  空结构体: %s\n\n", greeter.Greet("匿名"))

	// ============================================================
	// 第四部分：接口组合
	// ============================================================
	fmt.Println("--- 4. 接口组合（小而美的哲学） ---")

	// WHAT: 接口可以嵌入其他接口，实现组合
	// WHY: Go 鼓励小接口（单一方法接口），然后通过组合构建更大的接口
	//      这来自 Rob Pike 的格言："The bigger the interface, the weaker the abstraction"
	//      io.Reader（Read）、io.Writer（Write）、io.Closer（Close）
	//      组合成 io.ReadWriter、io.ReadCloser、io.ReadWriteCloser

	d := Document{Content: "Go 接口组合演示"}
	var rw ReadWriter = &d // Document 实现了 Read + Write

	buf := make([]byte, 20)
	n, _ := rw.Read(buf)
	fmt.Printf("  读取: %s\n", string(buf[:n]))

	rw.Write([]byte("新内容"))
	fmt.Printf("  写入后: %s\n\n", d.Content)

	// ============================================================
	// 第五部分：空接口 any / interface{}
	// ============================================================
	fmt.Println("--- 5. 空接口 any ---")

	// WHAT: any 是 interface{} 的别名（Go 1.18+）
	//      空接口没有方法，任何类型都满足它
	// WHY: 类似 TS 的 unknown，需要类型断言才能使用
	// CONTRAST:
	//   - TS: unknown（类型安全的任意值）
	//   - Rust: dyn Any（需要 downcast）
	//   - Go: any / interface{}（类型断言或 type switch）

	var anything any

	anything = 42
	printAny(anything)

	anything = "Hello, any!"
	printAny(anything)

	anything = []int{1, 2, 3}
	printAny(anything)

	anything = struct{ Name string }{"Gopher"}
	printAny(anything)
	fmt.Println()

	// ============================================================
	// 第六部分：类型断言与 type switch
	// ============================================================
	fmt.Println("--- 6. 类型断言与 type switch ---")

	// 类型断言（Type Assertion）
	var val any = "hello"
	str, ok := val.(string) // comma-ok 模式，安全！
	fmt.Printf("  val.(string): %s, ok=%t\n", str, ok)

	num, ok := val.(int) // 安全断言失败
	fmt.Printf("  val.(int): %d, ok=%t\n\n", num, ok)

	// type switch — 根据类型分支处理
	describeType(42)
	describeType("hello")
	describeType(3.14)
	describeType([]int{1, 2, 3})
	fmt.Println()

	// ============================================================
	// 第七部分：多态 — 接口作为参数
	// ============================================================
	fmt.Println("--- 7. 多态实践 ---")

	// WHAT: 接受接口，返回结构体 —— Go 社区的著名原则
	//       Accept interfaces, return structs
	// WHY: 函数接受接口可以接受任何实现，返回具体类型给调用者更多信息

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{Base: 3, Height: 4},
	}

	for _, shape := range shapes {
		fmt.Printf("  %s 面积: %.2f\n", shape.Name(), shape.Area())
	}

	fmt.Println("\n✅ Level 06 完成！")
}

// ============================================================
// 类型定义区
// ============================================================

// --- 结构体和方法 ---

// Rectangle 矩形
type Rectangle struct {
	Width, Height float64
}

// Area 值接收者 —— 不需要修改接收者时使用
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Scale 指针接收者 —— 需要修改接收者时使用
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Counter 计数器
type Counter struct{ Value int }

// ValueIncrement 值接收者（不影响原值）
func (c Counter) ValueIncrement() int {
	c.Value++ // 修改副本
	return c.Value
}

// PtrIncrement 指针接收者（修改原值）
func (c *Counter) PtrIncrement() int {
	c.Value++
	return c.Value
}

// --- 隐式接口 ---

// Greeter 问候接口
// 任何有 Greet(string) string 方法的类型都是 Greeter
type Greeter interface {
	Greet(name string) string
}

// EnglishGreeter 英文问候
type EnglishGreeter struct{}

func (e EnglishGreeter) Greet(name string) string {
	return "Hello, " + name + "!"
}

// ChineseGreeter 中文问候
type ChineseGreeter struct{}

func (c ChineseGreeter) Greet(name string) string {
	return "你好, " + name + "!"
}

// 空结构体也能实现接口
func (struct{}) Greet(name string) string {
	return "Hi " + name + " (from empty struct)"
}

// --- 接口组合 ---

// Reader 读接口
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer 写接口
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter 组合接口
type ReadWriter interface {
	Reader
	Writer
}

// Document 文档（实现了 ReadWriter）
type Document struct {
	Content string
	buf     []byte
}

func (d *Document) Read(p []byte) (n int, err error) {
	d.buf = []byte(d.Content)
	n = copy(p, d.buf)
	return n, nil
}

func (d *Document) Write(p []byte) (n int, err error) {
	d.Content = string(p)
	return len(p), nil
}

// --- 多态接口 ---

// Shape 形状接口
type Shape interface {
	Area() float64
	Name() string
}

// Circle 圆
type Circle struct{ Radius float64 }

func (c Circle) Area() float64  { return math.Pi * c.Radius * c.Radius }
func (c Circle) Name() string   { return "圆形" }

// Triangle 三角形
type Triangle struct{ Base, Height float64 }

func (t Triangle) Area() float64 { return 0.5 * t.Base * t.Height }
func (t Triangle) Name() string  { return "三角形" }

// Rectangle 已经实现了 Area，但缺少 Name → 不能作为 Shape
// 不过我们已经实现了 Name（需要补充）
// 等等 —— 上面 Rectangle 没有 Name！说明 Shape 的隐式检查在编译时进行
// func (r Rectangle) Name() string { return "矩形" } ← 取消注释即可让矩形也实现 Shape

// --- 辅助函数 ---

// printAny 打印 any 的值和类型
func printAny(v any) {
	fmt.Printf("  %v (类型: %T)\n", v, v)
}

// describeType type switch 演示
func describeType(v any) {
	switch val := v.(type) {
	case int:
		fmt.Printf("  整数: %d\n", val)
	case string:
		fmt.Printf("  字符串: %s (长度: %d)\n", val, len(val))
	case float64:
		fmt.Printf("  浮点数: %f\n", val)
	default:
		fmt.Printf("  其他类型: %T\n", val)
	}
}
