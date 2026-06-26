// WHAT: Bug 02 — 指针接收者方法无法通过值类型赋值给接口
// ERROR: 当接口方法使用指针接收者实现时，值类型不满足接口
//        但指针类型满足。反之则两者都满足。
//
// ============================================================
// 编译错误信息：
// cannot use MyData{} (value of type MyData) as DataHandler value
//   in variable declaration: MyData does not implement DataHandler
//   (method Process has pointer receiver)
// ============================================================
//
// 为什么会这样：
//   方法集规则：
//   值类型 T 的方法集：所有值接收者方法
//   指针类型 *T 的方法集：所有值接收者 + 所有指针接收者方法
//
//   所以：指针接收者实现的方法，只能用指针赋给接口。
//
// CONTRAST（与已知语言对比）：
//   - Rust: &self vs &mut self 在 trait 中显式声明
//   - TS:   无此概念（没有指针）
//   - Go:   值/指针接收者的方法集不同，影响接口实现

package main

import "fmt"

// DataHandler 数据处理接口
type DataHandler interface {
	Process(val int)
}

// MyData 数据
type MyData struct {
	Value int
}

// Process 指针接收者（MyData 不满足 DataHandler，*MyData 满足）
func (d *MyData) Process(val int) {
	d.Value = val * 2
}

func main() {
	fmt.Println("接口 + 值/指针接收者 演示\n")

	// 指针类型 → 满足接口 ✅
	var h1 DataHandler = &MyData{Value: 10}
	h1.Process(5)
	fmt.Printf("指针接收者: h1 = %+v\n", h1)

	// 值类型 → 不满足接口 ❌（编译错误）
	// var h2 DataHandler = MyData{Value: 10} // ← 编译错误
	// h2.Process(5)

	fmt.Println("注意: MyData{} 不满足 DataHandler，因为 Process 是指针接收者")
	fmt.Println("     &MyData{} 才满足")
}
