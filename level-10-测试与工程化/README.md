# Level 10：测试、工程化与工具链

> **通关标准**：能独立写出表驱动测试、基准测试，用 `go test`/`go vet`/`go fmt` 等工具，并能向 Rust/TS 开发者解释 Go 的测试哲学。

---

## 本关目标

- [x] 掌握 `testing` 包：`TestXxx`、表驱动测试、子测试
- [x] 理解 `benchmark`：`BenchmarkXxx` 和性能分析
- [x] 学会常用工具：`go fmt`、`go vet`、`go doc`
- [x] 掌握测试覆盖率：`go test -cover`
- [x] 理解 Go 的工程化哲学（代码风格统一）

---

## 核心概念速查

### 测试函数签名

```go
func TestXxx(t *testing.T) { ... }
func BenchmarkXxx(b *testing.B) { ... }
func ExampleXxx() { ... }
// Output:
// expected output
```

### 表驱动测试范式

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b, want int
    }{
        {"正数", 1, 2, 3},
        {"零", 0, 0, 0},
        {"负数", -1, -2, -3},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Errorf("%d + %d = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

---

## 与已有知识的对比表

| 维度 | Go | Rust | TS |
|------|-----|------|-----|
| **测试框架** | `testing` 标准库 | `cargo test` 内置 | jest/vitest/mocha（第三方） |
| **断言** | 手动 if + t.Error() | assert_eq! / assert! | expect(x).toBe(y) |
| **Mock** | 手动接口/结构体（或 gomock） | mockall 等第三方 | jest.fn() |
| **基准测试** | `testing.B` 内置 | `cargo bench` (nightly) | 需要第三方 |
| **代码格式** | `go fmt`（强制统一） | `rustfmt` | prettier（可选） |

---

## 运行命令

```bash
go run main.go
go test ./calc/ -v
go test ./calc/ -bench=. -benchmem
go test ./calc/ -cover
```

---

## 自检清单

- [ ] 能手写出表驱动测试模板
- [ ] 理解 `t.Run` 子测试的用途
- [ ] 能写出 benchmark 并解读结果
- [ ] 能独立修复 bugs/ 目录下的 3 个错误
- [ ] 能向 jest 用户解释 Go 的测试为什么不需要 assert 库
