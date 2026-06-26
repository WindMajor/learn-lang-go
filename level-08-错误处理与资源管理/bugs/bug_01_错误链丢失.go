// WHAT: Bug 01 — 错误包装时丢失错误链
// ERROR: 使用 fmt.Errorf("%v", err) 包装错误会丢失原始错误链
//        导致 errors.Is / errors.As 无法穿透到原始错误
//
// ============================================================
// 运行结果：
// errors.Is: false ← 期望是 true!
// ============================================================
//
// 为什么会这样：
//   %v 只把原始错误的字符串嵌入到新错误中，但 errors.Is 无法"看到"原始错误。
//   必须用 %w（wrap）来保持错误链。
//
//   规则：
//     fmt.Errorf("... %v", err)  → 错误链断裂！
//     fmt.Errorf("... %w", err)  → 错误链完好
//
// CONTRAST（与已知语言对比）：
//   - Rust: anyhow::Error::context() 自动保持链
//   - Go:   必须注意 %w vs %v——容易遗忘

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("--- 错误链丢失演示 ---")

	// BUG: 用 %v 包装 → 丢失错误链
	original := os.ErrNotExist
	brokenWrap := fmt.Errorf("文件不存在: %v", original) // ← BUG! 应该用 %w
	fmt.Printf("brokenWrap: %v\n", brokenWrap)
	fmt.Printf("errors.Is(brokenWrap, os.ErrNotExist) = %t ← BUG!\n\n",
		errors.Is(brokenWrap, os.ErrNotExist))

	// 修复: 用 %w 包装 → 保持错误链
	goodWrap := fmt.Errorf("文件不存在: %w", original) // ← %w 保持链
	fmt.Printf("goodWrap: %v\n", goodWrap)
	fmt.Printf("errors.Is(goodWrap, os.ErrNotExist) = %t ✅\n",
		errors.Is(goodWrap, os.ErrNotExist))
}
