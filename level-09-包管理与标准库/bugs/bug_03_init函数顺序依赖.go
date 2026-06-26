// WHAT: Bug 03 — init 函数执行顺序陷阱
// ERROR: 依赖 init 函数的执行顺序是脆弱的，跨包 init 顺序由依赖图决定
//        同包内不同文件的 init 按文件名排序（不直观！）
//
// ============================================================
// 为什么会这样：
//   跨包：Go 保证按依赖关系初始化（被依赖的先 init），这是确定的。
//   同包多文件：按文件名**字母序**执行 init，这也是确定的但非常脆弱！
//       重命名文件就改变了初始化顺序。
//
//   最佳实践：不要依赖 init 的执行顺序。如有复杂初始化，用显式函数。
//
// 如何修复：
//   1. init 中只放简单逻辑（注册、设置默认值）
//   2. 复杂初始化用显式 Init() 函数
//   3. 不要跨 init 互相依赖

package main

import "fmt"

var globalConfig map[string]string

func init() {
	fmt.Println("init 1: 初始化 globalConfig")
	globalConfig = make(map[string]string)
}

func init() {
	fmt.Println("init 2: 设置默认值")
	globalConfig["host"] = "localhost"
	globalConfig["port"] = "8080"
}

func main() {
	fmt.Printf("配置: %v\n", globalConfig)
	fmt.Printf("host: %s, port: %s\n", globalConfig["host"], globalConfig["port"])

	fmt.Println("\n注意：同一文件中的 init 按出现顺序执行")
	fmt.Println("     不同文件的 init 按文件名排序执行（脆弱！）")
	fmt.Println("     最佳实践：不要依赖 init 的执行顺序")
}
