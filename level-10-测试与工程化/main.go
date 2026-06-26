// WHAT: Level 10 主代码 — 测试与工程化
// WHY: 展示 Go 的测试工具链和工程化实践
//      本关的核心代码在 calc/ 包中，main.go 演示工具链使用

package main

import (
	"fmt"

	"github.com/user/go-basic-learn/level-10/calc"
)

func main() {
	fmt.Println("========== 测试、工程化与工具链 ==========\n")

	// ============================================================
	// 第一部分：Go 工具链速览
	// ============================================================
	fmt.Println("--- 1. Go 工具链 ---")
	fmt.Println("  go fmt     — 格式化代码（强制统一，无争议）")
	fmt.Println("  go vet     — 静态分析（检查可疑代码）")
	fmt.Println("  go test    — 运行测试")
	fmt.Println("  go build   — 编译")
	fmt.Println("  go run     — 编译并运行")
	fmt.Println("  go doc     — 查看文档")
	fmt.Println("  go mod tidy — 整理依赖\n")

	// ============================================================
	// 第二部分：使用自定义包
	// ============================================================
	fmt.Println("--- 2. 使用 calc 包 ---")

	result := calc.Add(10, 20)
	fmt.Printf("  calc.Add(10, 20) = %d\n", result)

	r2, err := calc.Divide(10, 3)
	if err != nil {
		fmt.Printf("  除法错误: %v\n", err)
	} else {
		fmt.Printf("  calc.Divide(10, 3) = %.2f\n", r2)
	}

	_, err = calc.Divide(10, 0)
	fmt.Printf("  calc.Divide(10, 0) → %v\n", err)
	fmt.Println()

	// ============================================================
	// 第三部分：表驱动测试概念
	// ============================================================
	fmt.Println("--- 3. 测试理念 ---")
	fmt.Println("  Go 的测试哲学：")
	fmt.Println("  • 测试文件和代码在同一包内（xxx_test.go）")
	fmt.Println("  • 测试函数名必须以 Test 开头")
	fmt.Println("  • 表驱动测试（table-driven tests）是 Go 社区惯用法")
	fmt.Println("  • 没有断言库（assert），用 if + t.Error()")
	fmt.Println("  • 基准测试（Benchmark）内置支持")
	fmt.Println()
	fmt.Println("  运行测试：")
	fmt.Println("    go test ./calc/ -v")
	fmt.Println("    go test ./calc/ -bench=. -benchmem")
	fmt.Println("    go test ./calc/ -cover")
	fmt.Println()

	fmt.Println("✅ Level 10 完成！")
	fmt.Println("👉 运行 go test ./calc/ -v 查看测试")
}
