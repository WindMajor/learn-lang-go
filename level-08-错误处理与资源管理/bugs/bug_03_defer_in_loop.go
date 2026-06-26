// WHAT: Bug 03 — 循环中的 defer 资源泄漏
// ERROR: 在循环中使用 defer，资源只在函数退出时释放，而不是每次迭代结束
//        导致循环中积累大量未释放资源
//
// ============================================================
// 运行结果：
// 打开 1000 个文件，但只到函数结束才关闭 → 可能耗尽文件描述符！
// ============================================================
//
// 为什么会这样：
//   defer 延迟到函数返回时才执行，不是块结束时执行。
//   在循环中用 defer 意味着所有资源堆积到函数结束才释放。
//
// CONTRAST（与已知语言对比）：
//   - Rust: Drop 在 {} 块结束时自动执行（和 Go 完全不同）
//   - C++:  析构函数在作用域结束时调用（RAII）
//   - Go:   defer 是函数级别的延迟！
//
// 如何修复：
//   1. 把循环体提取为单独函数（推荐）
//   2. 不用 defer，手动在循环体内释放

package main

import (
	"fmt"
	"os"
)

// BUG: 循环中 defer（资源堆积）
func buggyLoop() {
	fmt.Println("循环中 defer（资源堆积）:")
	for i := 0; i < 5; i++ {
		f, err := os.CreateTemp("", fmt.Sprintf("bug-loop-%d", i))
		if err != nil {
			continue
		}
		defer f.Close() // ← BUG! 仅函数结束时关闭！
		fmt.Fprintf(f, "data %d\n", i)
		// f 不会在这里关闭，而是堆积到函数结束！
	}
	fmt.Println("所有文件将在函数返回时关闭（可能已经打开太多）")
}

// 修复 1：提取为函数
func processOne(i int) error {
	f, err := os.CreateTemp("", fmt.Sprintf("fixed-loop-%d", i))
	if err != nil {
		return err
	}
	defer f.Close() // ← f 在 processOne 返回时关闭
	fmt.Fprintf(f, "data %d\n", i)
	return nil
}

func fixedLoop() {
	fmt.Println("提取为函数（每次迭代后释放）:")
	for i := 0; i < 5; i++ {
		if err := processOne(i); err != nil {
			fmt.Printf("  处理 %d 失败: %v\n", i, err)
		}
	}
	fmt.Println("每次迭代完成都释放了资源")
}

func main() {
	buggyLoop()
	fmt.Println()
	fixedLoop()
}
