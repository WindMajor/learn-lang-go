// WHAT: Level 09 主代码 — 包管理、模块与标准库
// WHY: 展示 go modules、可见性规则、核心标准库用法
// CONTRAST: npm 的 package.json vs go.mod, cargo 的 Cargo.toml vs go.mod

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("========== 包管理、模块与标准库 ==========\n")

	// ============================================================
	// 第一部分：go.mod 与可见性规则回顾
	// ============================================================
	fmt.Println("--- 1. go.mod 与可见性 ---")
	// WHAT: go.mod 定义了模块路径和 Go 版本
	//       首字母大小写控制包级符号的可见性
	// WHY: Go 选择用大小写而非 public/private 关键字
	//      - 减少关键词
	//      - 从名字就看出是否导出（读代码时省掉一跳）
	//      - 代价：不能通过改大小写来区分"相同名字不同可见性"
	fmt.Println("  模块名: github.com/user/go-basic-learn/level-09")
	fmt.Println("  可见性规则: 首字母大写=导出, 小写=私有")

	// 调用本包内的导出函数和私有函数
	fmt.Printf("  导出版本: %s\n", GetVersion())
	fmt.Printf("  私有计数: %d\n\n", internalCounter)

	// ============================================================
	// 第二部分：os 包 — 操作系统接口
	// ============================================================
	fmt.Println("--- 2. os 包 ---")

	// 环境变量
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = "(未设置)"
	}
	fmt.Printf("  GOPATH: %s\n", goPath)
	fmt.Printf("  当前用户: %s\n", os.Getenv("USER"))

	// 文件操作
	tmpDir := os.TempDir()
	fmt.Printf("  临时目录: %s\n", tmpDir)

	// 检查文件是否存在
	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("  go.mod 存在")
	}

	// 执行命令
	fmt.Printf("  PID: %d\n\n", os.Getpid())

	// ============================================================
	// 第三部分：io 包 — I/O 抽象
	// ============================================================
	fmt.Println("--- 3. io 包 ---")
	// WHAT: io.Reader 和 io.Writer 是 Go 的 I/O 抽象基石
	//       fmt.Fprintf, json.Decode, http.Get 都基于这两个接口
	// WHY: 统一的 I/O 接口让所有"可读/可写"的东西都能互操作

	// io.Reader 演示
	reader := strings.NewReader("Hello, io.Reader!")
	buf := make([]byte, 8)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("  读取 %d 字节: %s\n", n, string(buf[:n]))
	}
	fmt.Println()

	// ============================================================
	// 第四部分：encoding/json — JSON 编解码
	// ============================================================
	fmt.Println("--- 4. encoding/json ---")
	// WHAT: json.Marshal / json.Unmarshal 是最常用的序列化工具
	// WHY: 结构体 tag（`json:"field_name"`）是 Go 的编解码约定
	//      和 Rust 的 #[serde(rename = "field_name")] 作用相同

	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email,omitempty"` // omitempty: 空值不输出
	}

	// 编码：Go → JSON
	u := User{Name: "张三", Age: 30}
	data, _ := json.Marshal(u)
	fmt.Printf("  编码: %s\n", string(data))

	// 美化输出
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Printf("  美化:\n%s\n", string(pretty))

	// 解码：JSON → Go
	jsonStr := `{"name": "李四", "age": 25, "email": "lisi@example.com"}`
	var u2 User
	json.Unmarshal([]byte(jsonStr), &u2)
	fmt.Printf("  解码: %+v\n\n", u2)

	// ============================================================
	// 第五部分：net/http — HTTP 客户端
	// ============================================================
	fmt.Println("--- 5. net/http 客户端 ---")

	// WHAT: Go 的标准库自带生产级 HTTP 客户端和服务端
	// WHY: 不需要 Express/Axios/Reqwest —— 标准库就够了

	// 简单的 HTTP GET（演示用，不实际请求外部）
	fmt.Println("  HTTP 客户端就绪（标准库自带，无第三方依赖）")

	// 自定义 HTTP 客户端（带超时）
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	_ = client // 使用 client 执行请求
	fmt.Printf("  HTTP Client 配置: timeout=%s\n\n", client.Timeout)

	// ============================================================
	// 第六部分：time 包
	// ============================================================
	fmt.Println("--- 6. time 包 ---")

	now := time.Now()
	fmt.Printf("  当前时间: %s\n", now.Format("2006-01-02 15:04:05"))
	// WHAT: Go 的时间格式化模板是 2006-01-02 15:04:05（Mon Jan 2 15:04:05 MST 2006）
	// WHY: Go 诞生于 2006 年？这是个传说。实际是选取了一个容易记忆的参考时间点

	fmt.Printf("  Unix 时间戳: %d\n", now.Unix())
	fmt.Printf("  一年后: %s\n\n", now.AddDate(1, 0, 0).Format("2006-01-02"))

	// ============================================================
	// 第七部分：fmt 格式化补充
	// ============================================================
	fmt.Println("--- 7. fmt 格式化速查 ---")
	name, age := "Gopher", 14
	fmt.Printf("  %%s: %s | %%d: %d | %%f: %f | %%t: %t\n", name, age, 3.14, true)
	fmt.Printf("  %%v: %v | %%+v: %+v | %%#v: %#v | %%T: %T\n",
		struct{ X, Y int }{1, 2}, struct{ X, Y int }{1, 2},
		struct{ X, Y int }{1, 2}, struct{ X, Y int }{1, 2})
	fmt.Printf("  %%q: %q | %%p: %p\n\n", "hello", &age)

	fmt.Println("✅ Level 09 完成！")
}

// GetVersion 导出函数（首字母大写 → 其他包可访问）
func GetVersion() string {
	return "1.0.0"
}

// internalCounter 私有变量（首字母小写 → 仅本包可见）
var internalCounter = 0

// incrementCounter 私有函数
func incrementCounter() {
	internalCounter++
}

func init() {
	incrementCounter()
	incrementCounter()
}
