# Level 05：切片、映射与引用语义

> **通关标准**：能独立使用 slice/map，理解底层数组共享、append 扩容、map 非并发安全、nil slice vs 空 slice，并能向 Rust 开发者解释"Go 的 slice 为什么不是 Vec"。

---

## 本关目标

- [x] 理解数组与切片的区别
- [x] 掌握 slice 底层结构（ptr + len + cap）和共享底层数组
- [x] 精通 `make` / `append` / `copy` / 子切片
- [x] 理解 map（引用类型）及非并发安全
- [x] 区分 nil slice vs 空 slice vs nil map

---

## 核心概念速查

### Slice 底层结构

```
slice = struct {
    ptr  *Elem   // 指向底层数组的指针
    len  int     // 当前长度
    cap  int     // 容量
}
```

### 关键操作

```go
s := make([]int, 3, 5)   // len=3, cap=5
s = append(s, 1, 2, 3)   // 追加
s2 := s[1:3]             // 子切片（共享底层数组！）
s3 := make([]int, len(s))
copy(s3, s)              // 深拷贝
```

---

## 与已有知识的对比表

| 维度 | Go | Rust | TS | Python |
|------|-----|------|-----|--------|
| **动态数组** | slice（ptr+len+cap） | Vec（堆分配，所有权） | Array（引用计数） | list（引用） |
| **子切片** | 共享底层数组（陷阱！） | 用 `&vec[1..3]` 借用 | `slice()` 浅拷贝 | `list[1:3]` 浅拷贝 |
| **扩容** | append 自动扩容（倍增） | `vec.push()` 自动扩容 | `push()` 动态 | `append()` 动态 |
| **Map** | 引用类型，非并发安全 | HashMap（所有权） | Map（引用） | dict（引用） |
| **nil 语义** | nil slice 可 range/len | 不允许 None 替代 | `null`/`undefined` | `None` |

---

## 运行命令

```bash
go run main.go
go run playground.go
cd bugs && go run bug_01_切片共享底层数组.go && go run bug_02_map并发写入.go && go run bug_03_append返回新切片.go
```

---

## 自检清单

- [ ] 能手写 slice 的 make/append/copy/子切片操作
- [ ] 理解子切片共享底层数组的陷阱及如何避免
- [ ] 能独立修复 bugs/ 目录下的 3 个错误
- [ ] 能向 Rust 开发者解释：为什么 Go 的 slice 没有所有权概念
- [ ] 理解 map 为什么不是并发安全的，以及如何安全使用
