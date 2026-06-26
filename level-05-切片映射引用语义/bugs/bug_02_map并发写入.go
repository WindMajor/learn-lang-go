// WHAT: Bug 02 — Map 非并发安全
// ERROR: Go 的 map 不是并发安全的。多 goroutine 同时写 map 会导致
//        fatal error: concurrent map writes，程序直接崩溃
//
// ============================================================
// 运行结果（可能有多种）：
// fatal error: concurrent map writes
// 或者：fatal error: concurrent map read and map write
// 或者（极少数）：静默数据损坏
// ============================================================
//
// 为什么会这样：
//   Go 的 map 为单线程场景优化（无锁，速度快）
//   并发读写 map 时，Go runtime 会检测到冲突并直接 fatal（不是 panic）
//   这是不可恢复的错误（不像普通 panic 可以用 recover 捕获）
//
// CONTRAST（与已知语言对比）：
//   - Rust: HashMap 不并发安全但编译器阻止你共享（Send/Sync）
//   - TS:   单线程无此问题
//   - Go:   map 运行时检测并发冲突 → fatal error
//
//   关键：Go 提供了并发安全的选择：
//     1. sync.Mutex / sync.RWMutex 保护 map
//     2. sync.Map（适合特定场景）
//
// 如何修复：
//   var mu sync.Mutex
//   mu.Lock()
//   m[key] = value
//   mu.Unlock()

package main

import (
	"fmt"
	"sync"
)

func main() {
	// BUG：并发写 map（运行会崩溃）
	fmt.Println("并发写 map 演示（已被注释，取消注释运行看崩溃效果）")

	// 以下代码若取消注释会在运行时 fatal error
	// m := make(map[int]int)
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	//     wg.Add(1)
	//     go func(i int) {
	//         defer wg.Done()
	//         m[i] = i * 10 // ← 并发写！
	//     }(i)
	// }
	// wg.Wait()

	// ============================================================
	// 修复版本 1：用 sync.Mutex 保护
	// ============================================================
	fmt.Println("\n--- 修复 1：sync.Mutex ---")
	m1 := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m1[i] = i * 10
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("  m1: %v (安全)\n\n", m1)

	// ============================================================
	// 修复版本 2：用 sync.Map
	// ============================================================
	fmt.Println("--- 修复 2：sync.Map ---")
	var m2 sync.Map
	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func(i int) {
			defer wg2.Done()
			m2.Store(i, i*10)
		}(i)
	}
	wg2.Wait()

	fmt.Print("  m2: ")
	m2.Range(func(k, v interface{}) bool {
		fmt.Printf("%d:%d ", k, v)
		return true
	})
	fmt.Println("(安全)\n")

	fmt.Println("注意：sync.Map 适合 key 相对稳定（写少读多）的场景")
	fmt.Println("     大部分情况用 map + sync.Mutex 就够了")
}
