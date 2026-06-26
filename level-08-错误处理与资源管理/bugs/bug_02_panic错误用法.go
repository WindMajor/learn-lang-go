// WHAT: Bug 02 — panic 用于常规错误处理
// ERROR: 用 panic 处理可预期的错误（如：文件不存在、网络超时）
//        这违背了 Go 的设计哲学，且无法被调用者优雅处理
//
// ============================================================
// 为什么会这样：
//   Go 社区共识：panic 用于不可恢复的程序错误（bug），
//   如数组越界、nil 指针解引用。这些是"程序员犯错了"。
//
//   可预期的业务错误（网络失败、文件不存在、验证失败）应该用 error 返回。
//
// CONTRAST（与已知语言对比）：
//   - Rust: panic! → 不可恢复（等价 Go 的 panic）
//            Result<T, E> → 可恢复（等价 Go 的 error 返回值）
//   - TS:   throw → try-catch（都用于可恢复错误）
//   - Go:   error → 可恢复
//            panic → 不可恢复（但可用 recover）

package main

import (
	"errors"
	"fmt"
)

// BUG: 把 panic 当异常用
// func getUserBad(id int) string {
//     if id <= 0 {
//         panic("非法用户 ID") // ← BUG! 应该返回 error
//     }
//     return "User" + string(rune(id))
// }

// 正确: 用 error 返回
func getUserGood(id int) (string, error) {
	if id <= 0 {
		return "", errors.New("非法用户 ID")
	}
	return fmt.Sprintf("User%d", id), nil
}

func main() {
	// 正确使用方式
	name, err := getUserGood(0)
	if err != nil {
		fmt.Printf("获取用户失败: %v\n", err)
	} else {
		fmt.Printf("用户名: %s\n", name)
	}

	name, err = getUserGood(42)
	if err != nil {
		fmt.Printf("获取用户失败: %v\n", err)
	} else {
		fmt.Printf("用户名: %s\n", name)
	}

	fmt.Println("\n规则: panic 用于 Bug，error 用于业务错误。")
	fmt.Println("      Go 中没有 checked exception 或 Result 类型")
	fmt.Println("      你不检查 error，编译器不会提醒你 → 需要纪律！")
}
