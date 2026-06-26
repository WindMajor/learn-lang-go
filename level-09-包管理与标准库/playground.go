// WHAT: 沙盒文件 — 标准库实验
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("🎮 Level 09 沙盒\n")

	// ======== strings 实验 ========
	fmt.Println("--- strings ---")
	s := "Go 语言编程"
	fmt.Println("包含 Go?", strings.Contains(s, "Go"))
	fmt.Println("转大写:", strings.ToUpper(s))

	// ======== json 实验 ========
	fmt.Println("\n--- json ---")
	type Item struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	i := Item{ID: 1, Name: "项目"}
	data, _ := json.Marshal(i)
	fmt.Println("JSON:", string(data))

	// ======== 你的实验区 ========
	fmt.Println("\n--- 你的实验 ---")

	fmt.Println("✅ 沙盒完成！")
}
