// WHAT: Bug 01 — Switch 穿透陷阱
// ERROR: Go 的 switch 默认不穿透，但 C/TS 开发者可能习惯性依赖穿透
//        因此容易忘记 Go 的 case 自动 break
//
// ============================================================
// 运行结果：
// 分类: 单数（不会同时输出"在 1~3")
// ============================================================
//
// CONTRAST（与已知语言对比）：
//   - C:    switch(x){ case 1: doA(); // 穿透到 case 2! }
//   - TS:   同上（需要 break 阻止）
//   - Go:   switch 自动 break，需要穿透必须显式 fallthrough
//
// 如何修复：
//   如果需要多个 case 执行相同逻辑，用逗号分隔：
//     case 1, 2, 3: ...
//   或者用 fallthrough（极少用）

package main

import "fmt"

func main() {
	n := 1

	// BUG 演示：你以为会穿透到 case 2，实际不会
	switch n {
	case 1:
		fmt.Println("在 case 1")
		// 这里不会自动进入 case 2！必须用 fallthrough
	case 2:
		fmt.Println("在 case 2")
	case 3:
		fmt.Println("在 case 3")
	}

	// 如果要穿透：
	switch n {
	case 1:
		fmt.Print("case 1 → ")
		fallthrough // 显式穿透
	case 2:
		fmt.Println("case 2")
	}

	// 正确的多值匹配写法：
	switch n {
	case 1, 2, 3:
		fmt.Println("在 1~3 之间")
	case 4, 5:
		fmt.Println("在 4~5 之间")
	}
}
