// WHAT: Level 01 主代码 — Go 编译模型与类型系统
// WHY: 通过可运行代码展示 Go 的类型系统、零值哲学、编译流程
// CONTRAST: 作为 Rust/TS/C 用户，注意 Go 的"刻意简单"设计

// WHAT: package main 表示这是一个可执行程序（而非库）
// WHY: Go 要求入口必须是 main 包中的 main() 函数 —— 强制约定，不像 TS 可以从任意文件开始
// CONTRAST: Rust 也是 fn main()，C 是 int main()，TS 无此概念
package main

// WHAT: 导入标准库 fmt（格式化 I/O）和 reflect（运行时反射）
// WHY: Go 的导入是按包路径的，和你熟悉的 npm/cargo 不同：
//      - Go 用完整的模块路径（如 "fmt"）而不是相对路径
//      - 未使用的导入是编译错误（强迫清洁代码！）
// CONTRAST:
//   - TS: import { fmt } from "..."  —— 按符号导入
//   - Rust: use std::fmt;           —— 按模块导入
//   - Go: import "fmt"               —— 按包导入，所有符号共享一个命名空间
import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

// WHAT: main 函数 —— 程序的唯一入口
// WHY: 无参数、无返回值！程序退出状态通过 os.Exit(n) 设置
// CONTRAST: C 的 int main(int argc, char** argv)，Rust 的 fn main() -> Result
func main() {
	// ========================================
	// 第一部分：基本类型声明
	// ========================================
	fmt.Println("========== Go 基本类型系统 ==========")

	// WHAT: var 关键字声明变量，类型放在变量名后面
	// WHY: Go 的设计者认为 "var x int" 比 "int x" 更易读（从左到右读：声明 x 是 int）
	// CONTRAST:
	//   - C/C++:    int x = 42;
	//   - Rust:     let x: i32 = 42;
	//   - TS:       let x: number = 42;
	//   - Go:       var x int = 42    ← 类型后置！
	var a int = 42
	var b float64 = 3.14
	var c bool = true
	var d string = "Hello, Go!"
	var e byte = 'A'  // byte = uint8，Go 没有 char 类型！
	var f rune = '中'  // rune = int32，可以存任意 Unicode 码点

	// WHAT: %T 打印类型，%v 打印值
	fmt.Printf("a = %d, 类型: %T\n", a, a)
	fmt.Printf("b = %f, 类型: %T\n", b, b)
	fmt.Printf("c = %t, 类型: %T\n", c, c)
	fmt.Printf("d = %s, 类型: %T\n", d, d)
	fmt.Printf("e = %c (%d), 类型: %T\n", e, e, e)
	fmt.Printf("f = %c (%d), 类型: %T\n", f, f, f)

	// ========================================
	// 第二部分：零值哲学
	// ========================================
	fmt.Println("\n========== 零值初始化（Zero Value）==========")

	// WHAT: 声明但不初始化 → Go 自动赋零值
	// WHY: 这是 Go 的核心设计哲学 —— "让零值可用"（Make the zero value useful）
	//      sync.Mutex 零值就是可用的锁，bytes.Buffer 零值就是可写的缓冲区
	// CONTRAST:
	//   - C:    int x;         // x 是垃圾值！未定义行为！
	//   - Rust: let x: i32;    // 编译错误！必须初始化（除非 unsafe）
	//   - TS:   let x: number; // x === undefined（不是 number！）
	//   - Go:   var x int      // x == 0，确定可用
	var (
		zInt    int
		zFloat  float64
		zBool   bool
		zString string
		zByte   byte
		zRune   rune
	)

	fmt.Printf("int 零值:     %d\n", zInt)         // 0
	fmt.Printf("float64 零值: %f\n", zFloat)       // 0.000000
	fmt.Printf("bool 零值:    %t\n", zBool)        // false
	fmt.Printf("string 零值:  %q\n", zString)       // ""
	fmt.Printf("byte 零值:    %d\n", zByte)        // 0
	fmt.Printf("rune 零值:    %d\n", zRune)        // 0

	// WHAT: 重点 —— string 零值是 "" 不是 nil！
	// WHY: Go 的 string 是值类型（底层 struct{ptr, len}），不能为 nil
	//      只有 pointer/slice/map/chan/interface/function 可以为 nil
	// CONTRAST:
	//   - TS:  let s: string = null;  // 合法！（strictNullChecks 下要加 |null）
	//   - Rust: let s: Option<String> = None;  // 用 Option 表达空值
	//   - Go:  var s string           // s == ""，永远不会 nil

	// ========================================
	// 第三部分：类型推导（Type Inference）
	// ========================================
	fmt.Println("\n========== 类型推导（:=）==========")

	// WHAT: := 短变量声明 —— 声明 + 初始化 + 类型推导，一步完成
	// WHY: Go 的惯用法：函数内能用 := 就不用 var（除非需要零值）
	//      但包级别（全局）不能用 :=
	// CONTRAST:
	//   - Rust: let x = 42;     // 类型推导（let 万能）
	//   - TS:   const x = 42;   // 类型推导
	//   - Go:   x := 42         // 类型推导（仅函数内）
	name := "Gopher"          // 推导为 string
	age := 30                 // 推导为 int（不是 int32！注意 Go 的默认类型）
	pi := 3.14                // 推导为 float64（不是 float32！）
	isAwesome := true         // 推导为 bool

	fmt.Printf("name = %s, 类型: %T\n", name, name)
	fmt.Printf("age = %d, 类型: %T\n", age, age)
	fmt.Printf("pi = %f, 类型: %T\n", pi, pi)
	fmt.Printf("isAwesome = %t, 类型: %T\n", isAwesome, isAwesome)

	// WHAT: Go 的整数默认类型是 int（平台相关），浮点数默认是 float64
	// WHY: int 是为了和数组索引/长度匹配，float64 是为了精度
	// CONTRAST:
	//   - Rust: 整数字面量默认 i32，浮点数默认 f64
	//   - C:    42 是 int（平台相关），3.14 是 double
	//   - TS:   所有数字都是 number（IEEE 754 双精度）

	// ========================================
	// 第四部分：类型转换（必须显式！）
	// ========================================
	fmt.Println("\n========== 类型转换（必须显式！）==========")

	var x int = 10
	var y int64 = 20

	// WHAT: Go 中 int 和 int64 是完全不同的类型，不能直接运算
	// WHY: 防止隐式转换带来的精度丢失和 bug —— Go 宁愿多敲几个字，也不要隐式 Bug
	// CONTRAST:
	//   - C:    int x; long y; x + y;  // 自动提升，可能截断
	//   - Rust: let x: i32; let y: i64; x + y; // 编译错误！（也不能隐式转换）
	//   - Go:   int(10) + int64(20)    // 编译错误！必须显式 T(v)
	//   - TS:   let x: number;         // 不区分整型浮点
	z := x + int(y) // 必须显式转换 int64 → int
	fmt.Printf("x + int(y) = %d\n", z)

	var f1 float64 = 3.14
	f2 := int(f1) // 截断小数部分，不是四舍五入！
	fmt.Printf("float64 → int: %d（注意：截断，不是四舍五入）\n", f2)

	// ========================================
	// 第五部分：string 与 rune
	// ========================================
	fmt.Println("\n========== string 与 rune ==========")

	// WHAT: Go 的 string 是不可变的 byte 序列（UTF-8 编码）
	// WHY: Go 源码要求 UTF-8，string 存的就是 UTF-8 字节
	// CONTRAST:
	//   - Rust: String 是堆分配的 UTF-8 字符串，&str 是借用
	//   - C:    char* 是任意字节序列，编码由程序员保证
	//   - TS:   string 是 UTF-16（JavaScript 规范要求）
	s := "你好, Go!" // 11 个字节："你"(3) + "好"(3) + ", "(2) + "G"(1) + "o"(1) + "!"(1)

	fmt.Printf("字符串: %s\n", s)
	fmt.Printf("字节长度（len）: %d\n", len(s)) // 11（按字节算，不是按字符算！）
	fmt.Printf("字符数量（rune count）: %d\n", utf8.RuneCountInString(s)) // 7（按 Unicode 字符算）

	// WHAT: range 遍历 string 时按 rune（字符）遍历，不是按 byte
	fmt.Println("\n按 rune 遍历字符串：")
	for i, r := range s {
		fmt.Printf("  s[%d] = %c (U+%04X)\n", i, r, r)
	}

	// WHAT: 按索引取的是 byte，不是字符！
	fmt.Printf("\ns[0] = %d (byte)，不是字符！\n", s[0]) // 228，即 "你" 的 UTF-8 首字节

	// ========================================
	// 第六部分：nil 的概念
	// ========================================
	fmt.Println("\n========== nil 的概念 ==========")

	// WHAT: nil 是预定义标识符（零值），用于 pointer/slice/map/chan/func/interface
	// WHY: Go 的 nil 是有类型的，不同类型的 nil 是不同的（不像 C 的 NULL 就是 0）
	// CONTRAST:
	//   - C:    NULL 就是 (void*)0，所有指针类型通用
	//   - Rust: 无 nil！用 Option<T> 表达空值
	//   - TS:   null/undefined 是顶级类型，可赋值给任意类型（strict 模式除外）
	//   - Go:   nil 只能赋值给上述 6 种类型

	var ptr *int = nil
	var sli []int = nil
	var mp map[string]int = nil

	fmt.Printf("nil 指针: %v\n", ptr)
	fmt.Printf("nil 切片: %v, len=%d\n", sli, len(sli)) // 可以对 nil slice 取 len！返回 0
	fmt.Printf("nil map: %v, len=%d\n", mp, len(mp))     // 可以对 nil map 取 len！返回 0

	// WHAT: 对 nil map 读操作合法，但写操作会 panic！
	// _ = mp["key"]   // 合法，返回零值 0
	// mp["key"] = 1   // panic: assignment to entry in nil map

	// ========================================
	// 第七部分：类型断言与 type switch（预告）
	// ========================================
	fmt.Println("\n========== 类型反射（reflect 包预告）==========")

	// WHAT: reflect 可以在运行时获取类型信息
	// WHY: Go 是静态类型，但 reflect 提供了运行时内省能力（类似 RTTI）
	// CONTRAST:
	//   - TS: typeof x（编译时 + 运行时）
	//   - Rust: std::any::TypeId（有限反射）
	//   - Go: reflect.TypeOf（完整反射，但有性能开销）
	fmt.Printf("42 的动态类型: %v\n", reflect.TypeOf(42))
	fmt.Printf("3.14 的动态类型: %v\n", reflect.TypeOf(3.14))
	fmt.Printf("\"hello\" 的动态类型: %v\n", reflect.TypeOf("hello"))
	fmt.Printf("true 的动态类型: %v\n", reflect.TypeOf(true))

	fmt.Println("\n✅ Level 01 完成！你已经理解了 Go 的编译模型与类型系统。")
}
