# Level 04：函数、指针与结构体

> **通关标准**：能独立写出 Go 的函数（多返回值、命名返回值、变参）、defer 资源管理、指针（无运算）和结构体，并能向 C 开发者解释"Go 的指针为什么不能运算"。

---

## 本关目标

- [x] 掌握 Go 函数：多返回值、命名返回值、变参
- [x] 深刻理解 `defer`：LIFO 执行、参数求值时机、资源释放模式
- [x] 理解 Go 指针：有指针但无指针运算（与 C 的根本差异）
- [x] 掌握 `struct`：无 class，结构体是唯一的数据组合方式
- [x] 理解值类型 vs 引用类型的传参行为

---

## 核心概念速查

### 函数声明

```go
// 多返回值
func divide(a, b int) (int, error) { ... }

// 命名返回值
func split(sum int) (x, y int) { x = sum * 4 / 9; y = sum - x; return }

// 变参
func sum(nums ...int) int { ... }

// 函数是一等公民
var fn func(int) int = func(x int) int { return x * 2 }
```

### defer 关键规则

1. defer 参数在声明时求值（不是执行时！）
2. defer 按 LIFO（后进先出）顺序执行
3. defer 可以修改命名返回值

### 指针

```go
var p *int       // nil 指针
p = &x           // 取地址
*p = 10          // 解引用
// 没有 p++ ！没有指针运算！
```

---

## 与已有知识的对比表

| 维度 | Go | Rust | C | TS |
|------|-----|------|-----|-----|
| **多返回值** | 原生支持 | Result/Option（不完全是） | 通过指针参数模拟 | 数组/对象解构 |
| **defer** | ✅ 内置 | `Drop` trait + `defer!` 宏 | `atexit`（完全不同） | `try-finally` |
| **指针运算** | ❌ 不允许 | ❌ safe Rust 不允许 | ✅ `p++`, `p[n]` | N/A |
| **struct** | 无 class，struct 唯一 | struct（数据）+ impl（行为） | struct（纯数据） | class（数据+方法） |
| **变参** | `args ...int` | 宏（macro） | `va_list`（危险） | 展开运算符 `...` |

---

## 运行命令

```bash
go run main.go
go run playground.go
cd bugs && go run bug_01_defer求值时机.go && go run bug_02_命名返回值遮蔽.go && go run bug_03_指针接收者遗漏.go
```

---

## 自检清单

- [ ] 能手写出带 defer 资源清理的函数模板
- [ ] 理解 defer 的参数求值时机（不在执行时，在声明时）
- [ ] 能向 C 开发者解释：Go 的指针不能运算，这意味着什么
- [ ] 能独立修复 bugs/ 目录下的 3 个错误
- [ ] 理解命名返回值 + defer 的组合用法
