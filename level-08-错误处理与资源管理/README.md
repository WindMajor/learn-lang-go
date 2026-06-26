# Level 08：错误处理与资源管理

> **通关标准**：能独立写出 Go 风格的错误处理代码，理解 error 接口、errors.Is/As、defer 资源链式释放、panic/recover 的边界，并能向 Rust 开发者解释"为什么 Go 不用 Result"。

---

## 本关目标

- [x] 理解 Go 的 `error` 接口（`Error() string`）
- [x] 掌握错误包装：`fmt.Errorf("%w")`、`errors.Is`、`errors.As`
- [x] 深入 defer 资源链式释放模式
- [x] 理解 panic/recover 的正确使用场景（不是异常！）
- [x] 建立 Go 错误处理 vs Rust Result 的心智模型

---

## 核心概念速查

### error 接口

```go
type error interface {
    Error() string
}
```

### 错误包装

```go
err := fmt.Errorf("打开文件失败: %w", originalErr)
errors.Is(err, fs.ErrNotExist)  // 沿链查找
errors.As(err, &pathErr)        // 沿链类型匹配
```

### Recover 规则

- 只能在 defer 函数内部使用
- 捕获 panic 的值，防止程序崩溃
- 不应该用于常规错误处理！

---

## 与已有知识的对比表

| 维度 | Go | Rust | TS |
|------|-----|------|-----|
| **错误传递** | 返回值 `error`（手动 if err != nil） | `Result<T, E>`（? 运算符简写） | `throw` / `try-catch` |
| **错误链** | `%w` 包装 + `errors.Is/As` | `thiserror` / `anyhow` chain | `cause` 属性（非标准） |
| **panic vs 异常** | panic 是"不应该发生的错误" | `panic!` 是不可恢复的 | `throw` 可以 catch |
| **资源释放** | `defer`（显式） | `Drop` trait（自动） | `try-finally` |
| **强迫处理** | 不强迫！可以忽略 error 返回值 | 编译时强迫（Result 必须用） | 不强迫 |

---

## 运行命令

```bash
go run main.go
go run playground.go
cd bugs && go run bug_01_错误链丢失.go && go run bug_02_panic错误用法.go && go run bug_03_defer_in_loop.go
```

---

## 自检清单

- [ ] 能手写出自定义 error 类型和错误包装
- [ ] 理解 `errors.Is` vs `errors.As` 的差别
- [ ] 能向 Rust 开发者解释：Go 的 error 和 Rust 的 Result 谁更"安全"
- [ ] 知道 panic/recover 的正确使用边界
- [ ] 能独立修复 bugs/ 目录下的 3 个错误
