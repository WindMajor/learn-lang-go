# Level 03：控制流与 Go 惯用法

> **通关标准**：能独立写出 Go 风格的条件判断、循环和分支，理解 Go 为什么只有 `for` 一种循环、为什么 switch 不需要 break，并能对 Rust 开发者说清 match vs switch 的设计差异。

---

## 本关目标

- [x] 掌握 `if`（支持短声明）的 Go 惯用模式
- [x] 精通 `for`（Go 唯一循环，无 while/do-while）
- [x] 理解 `switch`（自动 break、无穿透、支持表达式和无表达式两种模式）
- [x] 掌握 `range`（遍历 array/slice/map/string/channel）
- [x] 接受 Go 没有三目运算符的现实（并理解原因）

---

## 核心概念速查

### for 的四种形态

```go
// 1. 标准 C 风格
for i := 0; i < 10; i++ { }

// 2. while 风格（Go 没有 while）
for condition { }

// 3. 无限循环
for { }

// 4. range 遍历
for i, v := range slice { }
```

### switch 的两种模式

```go
// 表达式模式
switch x {
case 1: ...
case 2: ...
default: ...
}

// 无表达式模式（替代 if-else 链）
switch {
case x < 0: ...
case x == 0: ...
default: ...
}
```

---

## 与已有知识的对比表

| 维度 | Go | Rust | C | TS |
|------|-----|------|-----|-----|
| **循环** | 只有 `for`（4 种形态） | `loop`/`while`/`for` | `for`/`while`/`do-while` | `for`/`while`/`do-while`/`for-of` |
| **switch 穿透** | **不穿透**（自动 break） | match 不穿透 | 需要 break（否则穿透） | switch 有穿透 |
| **三目运算符** | ❌ 没有！ | ❌ 用 if-else 或 match | `?:` | `?:` |
| **range** | 遍历多种数据结构 | `for x in iter` | 无（C17+ 后才慢慢有） | `for-of` |
| **分支表达力** | if/switch 是语句（无返回值） | if/match 是表达式（有返回值） | 语句 | 语句（但可用三目） |

---

## 运行命令

```bash
go run main.go
go run playground.go
cd bugs && go run bug_01_switch穿透.go && go run bug_02_for循环变量捕获.go && go run bug_03_range值拷贝陷阱.go
```

---

## 自检清单

- [ ] 能手写出 `for` 的四种形态（不含 IDE 提示）
- [ ] 理解为什么 Go 的 switch 不需要 break
- [ ] 能接受 Go 没有三目运算符（能用 if 或 switch 替代）
- [ ] 能独立修复 `bugs/` 目录下的 3 个错误
- [ ] 能向 Rust 开发者解释：Go 为什么没有 match 表达式
