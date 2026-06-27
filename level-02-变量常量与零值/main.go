// WHAT: Level 02 主代码 — Go 变量、常量与零值
// WHY: 深入展示 Go 的声明语法、iota 枚举、零值哲学、可见性规则
// CONTRAST: 对比 Rust 的 let/mut/const、C 的未定义行为、TS 的 undefined

package main

import (
	"fmt"
)

// ============================================================
// 包级别声明 —— 展示可见性规则和声明方式
// ============================================================

// WHAT: 首字母大写 → 导出（其他包可见），小写 → 私有（仅本包可见）
// WHY: Go 用大小写控制可见性，不需要 public/private 关键字
// CONTRAST:
//   - Rust: pub fn / fn
//   - TS:   export / (无)
//   - C:    extern（头文件声明）/ static（文件内可见）
//   - Go:   首字母大小写 —— 极简，但意味着你改不了名字来改变可见性

// MaxRetries 最大重试次数（导出，其他包可用）
const MaxRetries = 3

// defaultTimeout 默认超时（私有，仅本包可用）
const defaultTimeout = 30

// Version 当前版本（导出变量）
var Version = "1.0.0"

// buildCount 编译计数（私有变量）
var buildCount int

// ============================================================
// iota 常量生成器 —— Go 特色的无类型枚举
// ============================================================

// WHAT: iota 从 0 开始，每行自增 1
// WHY: Go 没有 enum 关键字，iota 是最轻量的枚举方案
// CONTRAST:
//   - Rust: enum Color { Red, Green, Blue }  ← 有类型安全
//   - TS:   enum Color { Red, Green, Blue }  ← 有反向映射
//   - Go:   iota                             ← 无类型安全（值就是 int），极简
const (
	StatusPending = iota // 0
	StatusRunning        // 1（隐式 = iota）
	StatusDone           // 2
	StatusFailed         // 3
)

// WHAT: iota 高级用法 —— 位掩码模式
// WHY: 配合位运算使用，是 Go 标准库中常见的模式（如 os.FileMode）
const (
	FlagRead  = 1 << iota // 1  (二进制 001)
	FlagWrite             // 2  (二进制 010)
	FlagExec              // 4  (二进制 100)
)

// WHAT: iota 跳过某个值
const (
	_         = iota // 跳过 0
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                   // 1 << 20 = 1048576
	GB                   // 1 << 30 = 1073741824
)

// WHAT: iota 多表达式（每行一组）
const (
	a, b = iota, iota + 1 // 0, 1
	c, d                  // 1, 2
	e, f                  // 2, 3
)

func main() {
	fmt.Println("========== Go 变量、常量与零值 ==========\n")

	// ============================================================
	// 第一部分：变量声明方式对比
	// ============================================================
	fmt.Println("--- 1. 变量声明方式 ---")

	// 方式 1：var 声明，零值初始化
	var count int
	var name string
	var active bool
	fmt.Printf("var count int      → %d (零值)\n", count)
	fmt.Printf("var name string    → %q (零值)\n", name)
	fmt.Printf("var active bool    → %t (零值)\n\n", active)

	// 方式 2：var 声明 + 初始化
	var maxSize int = 1024
	var appName = "learn-lang-go" // 类型推导
	fmt.Printf("var maxSize int = 1024  → %d\n", maxSize)
	fmt.Printf("var appName = \"...\"     → %s (类型: %T)\n\n", appName, appName)

	// 方式 3：:= 短声明（仅函数内可用）
	// WHY: 这是 Go 最惯用的局部变量声明方式
	//      你可以把它理解为 Rust 的 let（自动推导），但不能用于包级别
	pi := 3.14159
	message := "Hello, := world!"
	fmt.Printf("pi := 3.14159          → %f (类型: %T)\n", pi, pi)
	fmt.Printf("message := \"...\"       → %s\n\n", message)

	// 方式 4：多变量声明
	// CONTRAST: 这和 Rust 的 let (x, y) = (1, 2) 类似，和 TS 的解构类似
	x, y, z := 1, 2, 3
	firstName, lastName := "Alan", "Turing"
	fmt.Printf("x, y, z := 1, 2, 3     → %d, %d, %d\n", x, y, z)
	fmt.Printf("firstName, lastName := → %s %s\n\n", firstName, lastName)

	// ============================================================
	// 第二部分：零值哲学的工程价值
	// ============================================================
	fmt.Println("--- 2. 零值工程价值 ---")

	// WHAT: Go 的零值设计让很多类型"声明即用"
	// WHY: 这是深思熟虑的设计 —— 零值必须是"有用的默认值"
	//
	// 例 1：sync.Mutex 零值就是可用的锁（不需要 NewMutex()）
	// 例 2：bytes.Buffer 零值就是可写的缓冲区（不需要 malloc）
	// 例 3：int 零值是 0（计数器天然从 0 开始）
	// 例 4：bool 零值是 false（最安全的默认状态）

	// CONTRAST:
	//   C:    int count; // 垃圾值！必须显式 = 0
	//   Rust: let count: i32; // 编译错误！必须初始化
	//   Go:   var count int  // count == 0，直接用

	// 结构体零值演示
	type Config struct {
		Host    string
		Port    int
		UseSSL  bool
		Timeout float64
	}
	var cfg Config // 所有字段都是零值
	fmt.Printf("Config 零值: Host=%q, Port=%d, UseSSL=%t, Timeout=%f\n",
		cfg.Host, cfg.Port, cfg.UseSSL, cfg.Timeout)
	fmt.Printf("  注意：Port=0 可能不是合法端口，这说明零值也有局限性\n\n")

	// ============================================================
	// 第三部分：iota 全面演示
	// ============================================================
	fmt.Println("--- 3. iota 枚举 ---")
	fmt.Printf("StatusPending = %d\n", StatusPending)
	fmt.Printf("StatusRunning = %d\n", StatusRunning)
	fmt.Printf("StatusDone    = %d\n", StatusDone)
	fmt.Printf("StatusFailed  = %d\n\n", StatusFailed)

	fmt.Println("位掩码常量：")
	fmt.Printf("FlagRead  = %03b (%d)\n", FlagRead, FlagRead)
	fmt.Printf("FlagWrite = %03b (%d)\n", FlagWrite, FlagWrite)
	fmt.Printf("FlagExec  = %03b (%d)\n\n", FlagExec, FlagExec)

	// 位掩码组合（用 | 组合，用 & 检测）
	permission := FlagRead | FlagWrite
	fmt.Printf("FlagRead | FlagWrite = %03b (%d)\n", permission, permission)
	fmt.Printf("Has FlagRead?  %t\n", permission&FlagRead != 0)
	fmt.Printf("Has FlagExec?  %t\n\n", permission&FlagExec != 0)

	fmt.Println("文件大小常量：")
	fmt.Printf("KB = %d\n", KB)
	fmt.Printf("MB = %d\n", MB)
	fmt.Printf("GB = %d\n\n", GB)

	fmt.Printf("iota 多表达式: a=%d, b=%d, c=%d, d=%d, e=%d, f=%d\n\n", a, b, c, d, e, f)

	// ============================================================
	// 第四部分：可见性规则
	// ============================================================
	fmt.Println("--- 4. 可见性规则 ---")
	// WHAT: 导出符号首字母大写，私有符号首字母小写
	// WHY: 消除 public/private 关键字，代码更简洁
	//      但代价是：你不能把私有的改名为大写的"外表"而不改变含义
	fmt.Printf("导出常量 MaxRetries:  %d\n", MaxRetries)
	fmt.Printf("私有常量 defaultTimeout: %d (同包可访问)\n", defaultTimeout)
	fmt.Printf("导出变量 Version:     %s\n", Version)
	fmt.Printf("私有变量 buildCount:  %d\n\n", buildCount)

	// ============================================================
	// 第五部分：常量表达式
	// ============================================================
	fmt.Println("--- 5. 常量表达式 ---")

	// WHAT: Go 的 const 只能是编译期可计算的常量
	// WHY: const 在编译时求值，不允许运行时计算
	// CONTRAST:
	//   - Rust: const 也是编译期，和 Go 类似
	//   - TS:   const 只是不可变绑定，值可以是运行时计算的
	//   - C:    #define 是预处理宏（文本替换），const 是编译期
	const (
		SecondsPerMinute = 60
		MinutesPerHour   = 60
		SecondsPerHour   = SecondsPerMinute * MinutesPerHour // 编译期计算
	)
	fmt.Printf("1 小时 = %d 秒\n\n", SecondsPerHour)

	// WHAT: 无类型常量 —— Go 特有的概念
	// WHY: 无类型常量有更高精度，在赋值时才确定类型
	//      类似 Rust 的 {integer} 占位符但不需要 turbofish
	const BigNumber = 1 << 62              // 无类型常量，可赋值给 int64
	const PrecisePi = 3.141592653589793238 // 无类型浮点常量
	var i64 int64 = BigNumber
	var f32 float32 = PrecisePi
	fmt.Printf("无类型常量 BigNumber → int64: %d\n", i64)
	fmt.Printf("无类型常量 PrecisePi → float32: %f (精度丢失)\n", f32)
	// 如果 BigNumber 是 int64，无法赋给 int（因为 int 和 int64 不同）
	// 但无类型常量可以自适应！
	var asInt int = BigNumber // 自动适配！
	fmt.Printf("无类型常量 BigNumber → int: %d\n\n", asInt)

	// ============================================================
	// 第六部分：变量遮蔽（Shadowing）—— 常见陷阱
	// ============================================================
	fmt.Println("--- 6. 变量遮蔽 ---")

	outer := "外部变量"
	fmt.Printf("初始: outer = %s\n", outer)

	if true {
		outer := "内部新变量" // := 创建新变量，遮蔽外部变量！
		fmt.Printf("  块内: outer = %s\n", outer)
	}
	fmt.Printf("  块外: outer = %s (未被修改！)\n", outer)
	// CONTRAST:
	//   - Rust: 同样会遮蔽，但 Rust 程序员对此更习惯（let 本身就允许 shadowing）
	//   - Go:   := 看起来像赋值，容易误以为是修改而非新建

	// 正确做法：如果真想修改外部变量
	if true {
		outer = "真的修改了" // 用 = 而不是 :=
	}
	fmt.Printf("  修改后: outer = %s\n\n", outer)

	// ============================================================
	// 第七部分：init 函数（预告）
	// ============================================================
	fmt.Println("--- 7. init 函数 ---")
	fmt.Printf("init 函数已执行，buildCount = %d\n\n", buildCount)

	fmt.Println("✅ Level 02 完成！")
}

// WHAT: init 函数 —— 包初始化时自动调用
// WHY: 用于包级别的初始化工作（设置默认值、注册驱动等）
//      init 在 main 之前执行，顺序由依赖关系决定
// CONTRAST:
//   - Rust: 没有 init 函数，依赖 const/static + lazy_static/once_cell
//   - TS:   没有等价物，用模块顶层代码（有副作用）
//   - Python: __init__.py / 模块级代码
//   - Go:   init() 无参数无返回值，可定义多个（执行顺序不确定！）
func init() {
	buildCount = 42
	// 多个 init 函数按源文件名排序执行（但不应依赖此行为）
}
