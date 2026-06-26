# Level 07：并发：Goroutine 与 Channel

> **通关标准**：能独立写出 goroutine + channel + select 的并发程序，理解 GMP 调度模型，并能向 TS/Rust 开发者解释 Go 的 CSP 并发哲学。

---

## 本关目标

- [x] 掌握 `go` 关键字启动 goroutine
- [x] 理解 channel：无缓冲 vs 有缓冲、方向约束
- [x] 精通 `select` 多路复用
- [x] 理解 Goroutine 闭包变量捕获陷阱
- [x] 初步了解 `sync.WaitGroup` 和 `context`
- [x] 对比 CSP 模型（Go）vs async/await（TS）vs 线程（Rust）

---

## 核心概念速查

### Channel 类型

```go
ch := make(chan int)       // 无缓冲（同步）
ch := make(chan int, 10)   // 有缓冲（异步）
var ch chan<- int          // 只写 channel
var ch <-chan int          // 只读 channel
```

### select 多路复用

```go
select {
case v := <-ch1:
    // ch1 有数据
case ch2 <- v:
    // 写入 ch2
case <-time.After(1 * time.Second):
    // 超时
default:
    // 非阻塞
}
```

---

## 与已有知识的对比表

| 维度 | Go (CSP) | Rust (async) | TS (async/await) |
|------|----------|-------------|-------------------|
| **并发模型** | CSP（通信顺序进程）：goroutine+channel | async/await + 多线程 | async/await（单线程事件循环） |
| **并发单元** | goroutine（~2KB 栈，用户态调度） | async task / OS thread | Promise（微任务/宏任务） |
| **通信方式** | "用通信共享内存" | channel (tokio) / Arc<Mutex> | 回调/Promise/事件 |
| **调度器** | GMP（Go 运行时） | tokio/async-std 运行时 | 浏览器/Node 事件循环 |
| **阻塞** | goroutine 阻塞不阻塞线程 | .await 让出执行权 | await 让出事件循环 |

---

## 运行命令

```bash
go run main.go
go run playground.go
cd bugs && go run bug_01_goroutine闭包变量.go && go run bug_02_channel未关闭死锁.go && go run bug_03_select竞态条件.go
```

---

## 自检清单

- [ ] 能手写出 goroutine + channel + select 的经典模式
- [ ] 理解无缓冲 channel 和有缓冲 channel 的区别
- [ ] 能向 TS 开发者解释：为什么 Go 不需要 async/await关键字
- [ ] 理解 goroutine 闭包变量捕获陷阱及修复
- [ ] 能独立修复 bugs/ 目录下的 3 个错误
