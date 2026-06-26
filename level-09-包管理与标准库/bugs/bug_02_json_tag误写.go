// WHAT: Bug 02 — JSON struct tag 错误
// ERROR: JSON tag 拼写错误或格式不对，导致编解码行为不符合预期
//
// ============================================================
// 运行结果：
// 编码: {"wrong_name":""}  ← 字段名不符预期
// 解码后: Name = ""  ← 字段未正确解码
// ============================================================
//
// 为什么会这样：
//   struct tag 的格式是 `json:"field_name"`（反引号包围！）
//   常见的错误：
//   1. 用了普通引号而不是反引号
//   2. tag 中 key 拼写错误（jsonn、json 没冒号）
//   3. 忘了 omitempty 导致空值也被序列化
//
// 如何修复：
//   1. 确保用反引号 ` 包围 tag
//   2. tag 格式：`json:"name,omitempty"`

package main

import (
	"encoding/json"
	"fmt"
)

// BUG: json tag 拼写错误
type BadStruct struct {
	Field string `jsonn:"field"` // ← BUG! jsonn 不是 json
}

// 修复: 正确 tag
type GoodStruct struct {
	Field string `json:"field"`
}

func main() {
	// BUG 演示
	b := BadStruct{Field: "hello"}
	data, _ := json.Marshal(b)
	fmt.Printf("错误 tag 编码: %s ← 字段名丢了\n", string(data))

	// 修复演示
	g := GoodStruct{Field: "hello"}
	data, _ = json.Marshal(g)
	fmt.Printf("正确 tag 编码: %s ← 字段名正确\n", string(data))
}
