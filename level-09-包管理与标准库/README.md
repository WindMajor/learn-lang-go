# Level 09：包管理、模块与标准库

> **通关标准**：能独立创建 Go module、理解包导入规则和首字母大小写可见性、熟练使用 os/io/net/http/encoding/json 等标准库，并能向 npm/cargo 用户解释 go modules 的工作方式。

---

## 本关目标

- [x] 掌握 go modules（go.mod / go.sum / go mod tidy）
- [x] 理解包导入规则：同模块、外部模块、标准库
- [x] 深入可见性规则（首字母大小写）
- [x] 熟悉常用标准库：os, io, net/http, encoding/json, time
- [x] 理解 `init()` 函数的用途和陷阱

---

## 核心概念速查

### go modules 命令

```bash
go mod init github.com/user/project   # 初始化模块
go mod tidy                           # 整理依赖
go get example.com/pkg@v1.2.3         # 添加依赖
go mod download                       # 下载依赖
```

### 可见性规则

```go
// 首字母大写 → 导出（public）
func PublicFunc() {}
type PublicType struct{ PublicField int }

// 首字母小写 → 私有（private，包内可见）
func privateFunc() {}
type privateType struct{ privateField int }
```

---

## 与已有知识的对比表

| 维度 | Go (go modules) | Rust (cargo) | TS (npm) |
|------|----------------|--------------|----------|
| **依赖声明** | go.mod（手动/自动） | Cargo.toml | package.json |
| **锁文件** | go.sum | Cargo.lock | package-lock.json |
| **版本规则** | 语义导入 + 最小版本选择 | SemVer | SemVer |
| **可见性** | 首字母大小写（无关键字） | pub 关键字 | export 关键字 |
| **多文件** | 同一包内自动共享 | 显式 mod 声明 | import/require |

---

## 运行命令

```bash
go run main.go
go run playground.go
cd bugs && go build ./...  # 尝试编译看错误
```

---

## 自检清单

- [ ] 能手写 go.mod 初始化和 go mod tidy
- [ ] 理解 Go 的可见性规则（大写导出）并知道为什么这样设计
- [ ] 能写出基本的 HTTP 服务端和客户端
- [ ] 能独立修复 bugs/ 目录下的 3 个错误
- [ ] 能向 npm 用户解释 go modules 的版本管理
