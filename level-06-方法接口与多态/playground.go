// WHAT: 沙盒文件 — 接口与方法实验
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("🎮 Level 06 沙盒\n")

	// ======== 隐式接口实验 ========
	fmt.Println("--- 隐式接口 ---")
	type Speaker interface{ Speak() string }

	type Dog struct{}
	func (d Dog) Speak() string { return "Woof!" }

	type Cat struct{}
	func (c Cat) Speak() string { return "Meow!" }

	var s Speaker
	s = Dog{}
	fmt.Println(s.Speak())
	s = Cat{}
	fmt.Println(s.Speak())

	// ======== 接口组合 ========
	fmt.Println("\n--- 接口组合 ---")
	type Closer interface{ Close() error }
	type ReadCloser interface {
		Speaker
		Closer
	}

	// ======== 类型断言 ========
	fmt.Println("\n--- 类型断言 ---")
	var x any = "hello"
	if s, ok := x.(string); ok {
		fmt.Printf("是字符串: %s\n", s)
	}

	// ======== 你的实验区 ========
	fmt.Println("\n--- 你的实验 ---")

	fmt.Println("✅ 沙盒完成！")
}
