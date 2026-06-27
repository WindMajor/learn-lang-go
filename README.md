# learn-lang-go — Go 语言闯关式学习项目

> **写给谁？** 一位熟悉 TypeScript（深度）、Rust（深度）、C（熟练）、Python（了解）的全栈开发者。
> **不是什么？** 不是小白教程，不讲"什么是变量"。
> **是什么？** 一份"差异地图"——让你用已知语言作锚点，快速建立 Go 的心智模型。

---

## 环境准备

```bash
# 确认 Go 版本 ≥ 1.22
go version

# 创建项目目录（如果你是第一次接触本项目）
mkdir -p ~/Dev/Learn/learn-lang-go
cd ~/Dev/Learn/learn-lang-go
```

---

## 通关路线图（11 关，每关 30~90 分钟）

| 关卡 | 主题 | 核心收获 | 与已有知识的衔接点 |
|------|------|----------|-------------------|
| **01** | 编译模型与类型系统 | 理解 Go 的编译流程、静态类型、零值哲学 | C 的编译链接、Rust 的 cargo build、TS 的 tsc |
| **02** | 变量、常量与零值 | 掌握 `:=`、`const`、`iota`、零值设计 | Rust 的 let/mut、C 的未初始化陷阱、TS 的 undefined |
| **03** | 控制流与 Go 惯用法 | 精通 `if`/`for`/`switch`/`range`，理解 Go 为何如此设计 | Rust 的 match/loop、C 的 while、TS 的 for-of |
| **04** | 函数、指针与结构体 | 多返回值、defer（LIFO）、指针（有但无运算）、struct | Rust 的 Result 模式、C 的指针运算、TS 的 class |
| **05** | 切片、映射与引用语义 | 理解 slice 底层数组共享、map 非并发安全、nil slice | Rust 的 Vec 所有权、TS 的 Array、Python 的 list |
| **06** | 方法、接口与多态 | **隐式接口**（Go 最核心的设计！）、接口组合 | TS 的显式接口、Rust 的 trait（显式 impl） |
| **07** | 并发：Goroutine 与 Channel | `go`/`chan`/`select`，理解 GMP 调度模型 | TS 的 async/await、Rust 的 async 运行时、Python 的 asyncio |
| **08** | 错误处理与资源管理 | `error` 接口、`errors.Is`/`As`、panic/recover | Rust 的 Result/Option、TS 的 try-catch、Python 的异常 |
| **09** | 包管理、模块与标准库 | go modules、导入规则、首字母大小写可见性 | npm/yarn、cargo、pip |
| **10** | 测试、工程化与工具链 | `testing` 包、表驱动测试、`go test`/`go vet` | Rust 的 cargo test、TS 的 jest/vitest |
| **11** | 🎓 毕业设计：并发下载器 | 完整项目：goroutine+channel+接口+HTTP+JSON+测试 | 综合运用所有知识 |

---

## 如何使用本项目

### 每关的标准操作流程

```bash
# 1. 进入关卡目录
cd level-01-编译模型与类型系统

# 2. 阅读 README.md（路线图 + 对比表 + 自检清单）
cat README.md

# 3. 直接运行主代码
go run main.go

# 4. 打开 playground.go，随意修改、破坏、实验
#    这就是你的沙盒，改坏了不心疼

# 5. 进入 bugs/ 目录，读懂每个错误案例
cd bugs
go run bug_01_短声明作用域陷阱.go  # 看错误
# 尝试自己修复，然后对照注释里的修复方案

# 6. 回到 README.md，完成自检清单
```

### 学习原则

- **代码即教程**：`.go` 文件中有详尽的中文注释（WHAT/WHY/CONTRAST），不需要对着文档学。
- **错误是最好的老师**：每个 bug 文件都经过验证，真实可编译/运行，错误信息真实。
- **对比锚点**：每个新概念都显式对 TS/Rust/C/Python，帮你建立差异地图。
- **通关标准**：每关末尾的自检清单告诉你是否真正掌握了，而非"看完了"。

---

## 项目总览

```
learn-lang-go/
├── README.md                              # 你在这里
├── level-01-编译模型与类型系统/
│   ├── README.md
│   ├── go.mod
│   ├── main.go                            # 主代码（详注）
│   ├── playground.go                      # 沙盒文件
│   └── bugs/
│       ├── bug_01_短声明作用域陷阱.go
│       ├── bug_02_零值判断混淆.go
│       └── bug_03_类型推导与默认类型.go
├── level-02-变量常量与零值/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── playground.go
│   └── bugs/
│       ├── bug_01_const未赋值iota.go
│       ├── bug_02_零值误判为未初始化.go
│       └── bug_03_全局变量初始化顺序.go
├── level-03-控制流与惯用法/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── playground.go
│   └── bugs/
│       ├── bug_01_switch_case穿透.go
│       ├── bug_02_for循环变量捕获.go
│       └── bug_03_range值拷贝陷阱.go
├── level-04-函数指针结构体/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── playground.go
│   └── bugs/
│       ├── bug_01_defer求值时机.go
│       ├── bug_02_命名返回值遮蔽.go
│       └── bug_03_指针接收者遗漏.go
├── level-05-切片映射引用语义/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── playground.go
│   └── bugs/
│       ├── bug_01_切片共享底层数组.go
│       ├── bug_02_map并发写入.go
│       └── bug_03_append返回新切片.go
├── level-06-方法接口与多态/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── playground.go
│   └── bugs/
│       ├── bug_01_接口nil值陷阱.go
│       ├── bug_02_值接收者不可修改.go
│       └── bug_03_空接口类型断言.go
├── level-07-并发Goroutine与Channel/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── playground.go
│   └── bugs/
│       ├── bug_01_goroutine闭包变量.go
│       ├── bug_02_channel未关闭死锁.go
│       └── bug_03_select竞态条件.go
├── level-08-错误处理与资源管理/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── playground.go
│   └── bugs/
│       ├── bug_01_错误链丢失.go
│       ├── bug_02_panic不适当的用法.go
│       └── bug_03_defer_in_loop.go
├── level-09-包管理与标准库/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── playground.go
│   └── bugs/
│       ├── bug_01_循环导入.go
│       ├── bug_02_未导出字段访问.go
│       └── bug_03_init函数隐患.go
├── level-10-测试与工程化/
│   ├── README.md
│   ├── go.mod
│   ├── main.go
│   ├── calc/
│   │   ├── calc.go
│   │   └── calc_test.go
│   └── bugs/
│       ├── bug_01_表驱动测试误用.go
│       ├── bug_02_benchmark干扰.go
│       └── bug_03_测试辅助函数错误.go
└── level-11-毕业设计-并发下载器/
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── downloader/
    │   ├── downloader.go
    │   ├── downloader_test.go
    │   ├── types.go
    │   └── progress.go
    └── examples/
        └── urls.txt
```

---

## 开始闯关

```bash
# 准备好了？Let's Go!
cd level-01-编译模型与类型系统
go run main.go
```
