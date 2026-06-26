// WHAT: 沙盒文件 — 错误处理实验
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("🎮 Level 08 沙盒\n")

	// ======== error 实验 ========
	fmt.Println("--- error ---")
	err := someFunc()
	if err != nil {
		fmt.Println("错误:", err)
	}

	// ======== errors.Is 实验 ========
	fmt.Println("\n--- errors.Is ---")
	base := errors.New("基础错误")
	wrapped := fmt.Errorf("包装: %w", base)
	fmt.Println("Is base?", errors.Is(wrapped, base))

	// ======== panic recover 实验 ========
	fmt.Println("\n--- panic recover ---")
	safeCall()

	// ======== 你的实验区 ========
	fmt.Println("\n--- 你的实验 ---")

	fmt.Println("✅ 沙盒完成！")
}

func someFunc() error {
	return errors.New("演示错误")
}
func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic:", r)
		}
	}()
	// panic("演示panic") // 取消注释看效果
	fmt.Println("正常执行")
}
