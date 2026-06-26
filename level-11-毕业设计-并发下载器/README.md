# Level 11 🎓 毕业设计：并发下载器

> **通关标准**：能独立完成此项目，并理解其中用到的所有 Go 特性：goroutine、channel、接口、错误处理、HTTP 客户端、JSON、测试、`go build` 打包。

---

## 项目描述

一个命令行并发文件下载器，支持：
- 多个 URL 并发下载（可配置并发数）
- 实时进度显示（速度、百分比、ETA）
- 支持从文件读取 URL 列表
- 错误重试与超时控制
- 下载结果 JSON 报告

### 用到的 Go 特性

| 特性 | 使用位置 |
|------|---------|
| **goroutine + channel** | 工作池并发下载 |
| **隐式接口** | 进度报告接口、存储接口 |
| **错误处理** | errors.Is/As、错误包装 |
| **defer** | 文件关闭、资源清理 |
| **struct 与方法** | 任务定义、进度追踪 |
| **net/http** | HTTP 下载 |
| **encoding/json** | 下载报告 |
| **testing** | 单元测试 + 表驱动测试 |
| **go build** | 编译为可执行文件 |

---

## 构建和运行

```bash
# 编译
go build -o go-downloader .

# 运行 — 下载单个文件
./go-downloader -url https://example.com/file.zip -o output/

# 运行 — 并发下载多个文件
./go-downloader -urls urls.txt -concurrency 5 -o ./downloads/

# 查看帮助
./go-downloader -h

# 直接运行（不编译）
go run main.go -url https://go.dev/dl/go1.22.0.src.tar.gz -o /tmp/dl/

# 运行测试
go test ./downloader/ -v

# 运行测试 + 覆盖率
go test ./downloader/ -cover
```

---

## urls.txt 格式示例

```
https://example.com/file1.zip
https://example.com/file2.pdf
https://example.com/image.jpg
```

---

## 项目结构

```
level-11-毕业设计-并发下载器/
├── README.md
├── go.mod
├── main.go                    # 入口，参数解析，编排
└── downloader/
    ├── types.go               # 类型定义与接口
    ├── downloader.go          # 核心下载逻辑
    ├── progress.go            # 进度追踪
    └── downloader_test.go     # 单元测试
```

---

## 自检清单

- [ ] 能独立运行此项目并下载文件
- [ ] 能向别人解释此项目中 goroutine 工作池的设计
- [ ] 理解 `select` 在此项目中如何用于超时和取消
- [ ] 能独立写出 downloader 包的单元测试
- [ ] 能用 `go build` 打包为可执行文件给同事使用
