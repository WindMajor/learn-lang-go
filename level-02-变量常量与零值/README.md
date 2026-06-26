# Level 02：变量、常量与零值

> **通关标准**：能独立写出 Go 的变量声明、常量定义、iota 枚举、理解零值哲学和可见性规则，并能向 C 开发者解释"为什么 Go 没有未定义行为"。

---

## 本关目标

- [x] 掌握 `var` / `:=` / `const` / `iota` 的用法和选择时机
- [x] 深入理解 Go 零值哲学的工程价值
- [x] 理解包级变量与局部变量的初始化顺序
- [x] 掌握 Go 可见性规则（首字母大小写决定导出）
- [x] 理解 Go 没有"未初始化"概念（与 C/Rust/TS 的对比）

---

## 核心概念速查

### 变量声明方式对比

| 方式 | 场景 | 示例 |
|------|------|------|
| `var x int` | 需要零值 | `var count int` （count=0） |
| `var x = val` | 包级别推导 | `var MaxSize = 1024` |
| `x := val` | 函数内短声明 | `name := "Gopher"` |
| `var x, y int` | 多变量同类型 | `var w, h int` |
| `var x, y = a, b` | 多变量推导 | `var name, age = "Tom", 30` |

### iota 常量生成器

```go
const (
    A = iota  // 0
    B         // 1（隐式重复上一行表达式）
    C         // 2
)
```

### 零值速查

| 类型 | 零值 | 是否可为 nil |
|------|------|-------------|
| `int`, `float64` | `0`, `0.0` | ❌ |
| `bool` | `false` | ❌ |
| `string` | `""` | ❌ |
| `struct` | 所有字段零值 | ❌ |
| `*T` (指针) | `nil` | ✅ |
| `[]T` (切片) | `nil` | ✅ |
| `map[K]V` | `nil` | ✅ |
| `chan T` | `nil` | ✅ |
| `func` | `nil` | ✅ |
| `interface` | `nil` | ✅ |

---

## 与已有知识的对比表

| 维度 | Go | Rust | C | TS |
|------|-----|------|-----|-----|
| **未初始化行为** | 强制零值 | 编译错误 | **未定义行为！** | `undefined` |
| **常量** | `const`（仅数字/字符串/布尔，编译期） | `const`（编译期） | `#define` / `const` | `const`（运行时） |
| **枚举** | `iota`（轻量，无类型安全） | `enum`（有类型安全） | `enum`（C 风格） | `enum` / union |
| **变量遮蔽** | 允许（新作用域） | 允许（推荐） | 允许（警告） | 允许（严格模式警告） |
| **可见性** | 首字母大小写！ | `pub` 关键字 | 头文件 + `extern` | `export` / 默认导出控制 |

---

## 运行命令

```bash
go run main.go
go run playground.go
cd bugs && for f in *.go; do echo "=== $f ===" && go run "$f"; done
```

---

## 自检清单

- [ ] 能手写出 `var` / `:=` / `const` / `iota` 的声明语法
- [ ] 能独立修复 `bugs/` 目录下的 3 个错误案例
- [ ] 能向 C 开发者解释：Go 零值初始化避免了哪些 C 语言陷阱
- [ ] 理解为什么 Go 选择首字母大小写控制可见性（而非 public/private 关键字）
- [ ] 能说出 `iota` 和 Rust enum 的差异
