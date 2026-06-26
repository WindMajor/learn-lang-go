# Level 06：方法、接口与多态

> **通关标准**：能独立设计 Go 的隐式接口，理解接口组合、空接口、类型断言，并能向 TS 开发者解释"为什么 Go 不需要 implements 关键字"。

---

## 本关目标

- [x] 掌握方法接收者（值接收者 vs 指针接收者）的正确选择
- [x] **深入理解 Go 的隐式接口**（Go 最核心的设计！）
- [x] 掌握接口组合（小而美的接口哲学）
- [x] 理解空接口 `any` / `interface{}` 和类型断言
- [x] 掌握接口 nil 值陷阱（最著名的 Go 陷阱之一）

---

## 核心概念速查

### 隐式接口（与 TS/Rust 的根本差异）

```go
// 定义接口
type Writer interface {
    Write([]byte) (int, error)
}

// 任何有 Write 方法的类型自动实现了 Writer
// 不需要写 implements 关键字！
type MyWriter struct{}
func (m MyWriter) Write(p []byte) (int, error) { ... }

var w Writer = MyWriter{} // ✅ 隐式满足接口
```

### 接口组合

```go
type Reader interface {
    Read([]byte) (int, error)
}
type ReadWriter interface {
    Reader
    Writer
}
```

---

## 与已有知识的对比表

| 维度 | Go | TypeScript | Rust |
|------|-----|------------|------|
| **接口实现** | **隐式**（无需 `implements`） | 显式 `implements` / 结构化类型 | 显式 `impl Trait for Type` |
| **接口设计** | 小而美（单一方法接口） | 可选属性、索引签名 | trait 可以有默认实现 |
| **多态** | 接口 + 类型断言 | 接口 + instanceof | trait object + dyn |
| **空接口** | `any` / `interface{}` | `unknown` | `dyn Any` |
| **方法定义** | `func (r Receiver) Method()` | class method / 对象方法 | `fn method(&self)` in impl |

---

## 运行命令

```bash
go run main.go
go run playground.go
cd bugs && go run bug_01_接口nil值陷阱.go && go run bug_02_值接收者不能改.go && go run bug_03_空接口类型断言.go
```

---

## 自检清单

- [ ] 能手写一个接口定义和两种实现（值接收者 + 指针接收者）
- [ ] 能向 TS 开发者解释：隐式接口和结构化类型的区别与联系
- [ ] 理解接口 nil 值陷阱：`var w io.Writer = (*MyWriter)(nil)` 中 w 不是 nil
- [ ] 能独立修复 bugs/ 目录下的 3 个错误
- [ ] 理解"接受接口，返回结构体"的 Go 设计原则
