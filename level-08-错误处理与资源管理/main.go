// WHAT: Level 08 主代码 — 错误处理与资源管理
// WHY: Go 的错误处理是显式、手动、不加糖的——理解这一哲学是关键
// CONTRAST: Rust 的 Result/?, JS 的 try-catch, Python 的异常

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("========== 错误处理与资源管理 ==========\n")

	// ============================================================
	// 第一部分：error 接口
	// ============================================================
	fmt.Println("--- 1. error 接口 ---")

	// WHAT: error 是 Go 内置的接口，只有一个方法 Error() string
	// WHY: 极简设计 —— 任何能描述自己的字符串的类型都是错误
	// CONTRAST:
	//   - Rust: trait Error: Debug + Display {}
	//   - TS:   Error class (有 message, stack, name 属性)
	//   - Go:   interface { Error() string } ← 极简到极致

	err1 := errors.New("这是一个基本错误")
	fmt.Printf("  errors.New: %v\n", err1)

	err2 := fmt.Errorf("格式化错误: %s", "详情")
	fmt.Printf("  fmt.Errorf: %v\n\n", err2)

	// ============================================================
	// 第二部分：自定义错误类型
	// ============================================================
	fmt.Println("--- 2. 自定义错误类型 ---")

	// WHAT: 任何实现 Error() string 的类型都是 error
	//       你可以在自定义错误中添加更多字段和信息
	valErr := &ValidationError{
		Field: "Email",
		Value: "invalid-email",
		Msg:   "格式不正确",
	}
	fmt.Printf("  自定义错误: %v\n", valErr)

	// 类型断言获取详细信息
	var ve *ValidationError
	if errors.As(valErr, &ve) {
		fmt.Printf("  错误详情: Field=%s, Value=%s\n\n", ve.Field, ve.Value)
	}

	// ============================================================
	// 第三部分：Go 惯用的错误处理模式
	// ============================================================
	fmt.Println("--- 3. 惯用错误处理 ---")

	// WHAT: 经典三行模式 (open → check error → defer close)
	// WHY: 显式、直接、不需要魔法 —— 但也很啰嗦
	// CONTRAST:
	//   - Rust: let f = File::open("x.txt")?; // 一行搞定
	//   - TS:   const f = await fs.open("x.txt"); // 可能抛异常
	//   - Go:   3 行（open, if err, defer close）

	filename := "/tmp/go-learn-demo.txt"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("  创建文件失败: %v\n", err)
	} else {
		defer f.Close()
		f.WriteString("Go 错误处理演示\n")
		fmt.Printf("  ✅ 已写入: %s\n", filename)
	}
	fmt.Println()

	// ============================================================
	// 第四部分：错误包装与 errors.Is / errors.As
	// ============================================================
	fmt.Println("--- 4. 错误包装 (Error Wrapping) ---")

	// WHAT: fmt.Errorf("%w", err) 包装错误，保留原始错误链
	// WHY: 让你能添加上下文信息，同时保留原始错误的可判断性
	//      errors.Is 检查错误链中是否存在特定错误
	//      errors.As 检查错误链中是否存在特定类型的错误
	// CONTRAST:
	//   - Rust: anyhow::Context 的 .context() / .with_context()
	//   - Go:   fmt.Errorf("%w") 包装 + errors.Is/As 查询

	originalErr := os.ErrNotExist
	wrappedErr := fmt.Errorf("读取配置文件失败: %w", originalErr)

	fmt.Printf("  原始错误: %v\n", originalErr)
	fmt.Printf("  包装错误: %v\n", wrappedErr)

	// errors.Is — 判断错误链中是否包含特定错误
	if errors.Is(wrappedErr, os.ErrNotExist) {
		fmt.Println("  ✅ errors.Is 找到 os.ErrNotExist!")
	}

	// errors.As — 提取特定类型的错误
	var pathErr *os.PathError
	if errors.As(wrappedErr, &pathErr) {
		fmt.Printf("  ✅ errors.As 提取 PathError: %s\n", pathErr.Path)
	} else {
		fmt.Println("  包装的不是 PathError")
	}
	fmt.Println()

	// ============================================================
	// 第五部分：panic 与 recover
	// ============================================================
	fmt.Println("--- 5. panic 与 recover ---")
	// WHAT: panic 是程序崩溃（类似 Rust 的 panic!，C 的 abort）
	//       recover 只能在 defer 中捕获 panic
	// WHY: panic 用于"不该发生的严重错误"，不是异常处理机制！
	//      普通错误应该用 error 返回值，不要用 panic
	// CONTRAST:
	//   - Rust: panic! → 不可恢复（大部分情况下），catch_unwind 可以捕获
	//   - TS:   throw → try-catch 捕获（异常是常规错误处理手段）
	//   - Go:   panic → recover（不应用于常规错误处理！）
	//   Go 和 Rust 的态度一致：panic 是"世界末日"，不要滥用

	// defer + recover 捕获 panic
	safeDiv := func(a, b int) (result int, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic 被捕获: %v", r)
			}
		}()
		result = a / b // b=0 时 panic
		return
	}

	r, err := safeDiv(10, 2)
	fmt.Printf("  safeDiv(10, 2) = %d, err=%v\n", r, err)

	r, err = safeDiv(10, 0)
	fmt.Printf("  safeDiv(10, 0) = %d, err=%v\n\n", r, err)

	// ============================================================
	// 第六部分：defer 链式资源释放
	// ============================================================
	fmt.Println("--- 6. defer 链式资源释放 ---")

	// WHAT: 多个 defer 可以串成资源释放链
	//       每个资源紧挨创建，defer 释放，形成清晰的模式
	fmt.Println("  资源链演示：")
	chainDemo()
	fmt.Println()

	fmt.Println("✅ Level 08 完成！")
}

// ValidationError 自定义验证错误
type ValidationError struct {
	Field string
	Value string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("验证错误 [%s=%q]: %s", e.Field, e.Value, e.Msg)
}

// chainDemo defer 链式资源释放
func chainDemo() {
	f1, err := os.CreateTemp("", "chain-demo-1")
	if err != nil {
		fmt.Printf("    打开文件1失败: %v\n", err)
		return
	}
	defer func() {
		f1.Close()
		fmt.Println("    释放 文件1")
	}()

	f2, err := os.CreateTemp("", "chain-demo-2")
	if err != nil {
		fmt.Printf("    打开文件2失败: %v\n", err)
		return
	}
	defer func() {
		f2.Close()
		fmt.Println("    释放 文件2")
	}()

	f1.WriteString("data1")
	f2.WriteString("data2")
	fmt.Println("    资源正在使用中...")
	// 函数返回时，释放顺序：f2 → f1（LIFO）
}
