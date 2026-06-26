# Level 01：编译模型与类型系统

> **通关标准**：能独立写出一个 Go 文件，用 `go run` 运行，理解 Go 的编译模型、静态类型、零值哲学，并能向 Rust/TS 开发者解释清楚"Go 编译与你们有何不同"。

---

## 本关目标

- [x] 理解 Go 的编译流程（`go build` / `go run` / `go install`）
- [x] 掌握 Go 的基本数据类型（int、float、string、bool、byte、rune）
- [x] 理解 Go 的**零值初始化**哲学（与 C 的未定义行为对照）
- [x] 学会类型推导（`:=` 的自动类型推导规则）
- [x] 体会 Go 的"简单"设计哲学

---

## 核心概念速查

### Go 的编译模型

```
┌──────────┐    ┌──────────┐    ┌──────────────┐
│  .go 源码 │───▶│  编译器   │───▶│ 静态链接二进制 │
└──────────┘    └──────────┘    └──────────────┘
                    │
                    ▼
            单文件也可编译
            无需头文件
            无需 Makefile
```

### 基本数据类型

| 类型 | 说明 | 零值 | 大小（64位系统） |
|------|------|------|-----------------|
| `bool` | 布尔 | `false` | 1 字节 |
| `int8/16/32/64` | 有符号整数 | `0` | 1/2/4/8 字节 |
| `uint8/16/32/64` | 无符号整数 | `0` | 1/2/4/8 字节 |
| `float32/64` | 浮点数 | `0.0` | 4/8 字节 |
| `complex64/128` | 复数 | `0+0i` | 8/16 字节 |
| `string` | 字符串 | `""` | 16 字节（ptr+len） |
| `byte` | uint8 别名 | `0` | 1 字节 |
| `rune` | int32 别名（Unicode 码点） | `0` | 4 字节 |
| `int/uint` | 平台相关（64位系统=64位） | `0` | 4/8 字节 |

### 零值哲学（Zero Value）

Go **强制**零值初始化，所有变量声明后都有确定的值：

```go
var i int       // i == 0（不是垃圾值！）
var b bool      // b == false
var s string    // s == ""（不是 nil！Go 的 string 不能为 nil）
```

---

## 与已有知识的对比表

| 维度 | Go | TypeScript | Rust | C | Python |
|------|-----|------------|------|-----|--------|
| **编译方式** | 编译为静态链接的 native 二进制 | tsc 转译为 JS（解释执行） | cargo build → native 二进制 | gcc/clang → native 二进制 | 解释执行（无编译） |
| **类型系统** | 静态类型 + 类型推导 | 结构化类型 + 类型推导 | 静态类型 + Hindley-Milner 推导 | 静态类型（弱类型检查） | 动态类型（运行时检查） |
| **变量初始化** | **零值初始化**（强制执行） | `undefined`（无默认值） | 必须显式初始化或使用 `Default` | **未定义行为**（垃圾值） | `NameError`（未定义则报错） |
| **字符串** | 不可变 byte 切片，非 nil | 不可变，可为 null/undefined | `&str` 引用 / `String` 堆分配 | `char*` 指针，手动管理 | 不可变，引用计数 |
| **nil/null** | 指针/接口/map/slice/chan 可为 nil | `null` / `undefined` | `None`（Option<T>） | `NULL`（宏，实际为 0） | `None` |
| **包管理** | go modules（go.mod） | npm/yarn（package.json） | cargo（Cargo.toml） | 无（手动头文件+链接） | pip（requirements.txt） |
| **入口点** | `main` 包中 `main()` 函数 | 无（脚本入口） | `fn main()` | `int main()` | 脚本顶层代码 |

---

## 运行命令

```bash
# 直接运行
go run main.go

# 编译并运行
go build -o myapp && ./myapp

# 查看 Go 环境
go env

# 格式化代码
go fmt main.go
```

---

## 自检清单

- [ ] 能手写出本关核心代码（类型声明、类型推导、`fmt.Printf` 格式化），不借助 IDE 提示
- [ ] 能独立修复 `bugs/` 目录下的 3 个错误案例
- [ ] 能向一个不懂 Go 的 Rust 开发者解释清楚：Go 的零值初始化 vs Rust 的 `Default` trait vs C 的未定义行为
- [ ] 能说出 `:=` 和 `var` 的使用场景区别
- [ ] 理解 `int` 和 `int64` 在 Go 中是完全不同的类型（不像 C 的隐式转换）
