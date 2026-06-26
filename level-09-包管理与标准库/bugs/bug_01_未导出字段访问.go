// WHAT: Bug 01 — 跨包访问未导出（小写）的字段/方法
// ERROR: 小写字母开头的标识符仅供同一包内访问
//        跨包访问会编译错误
//
// ============================================================
// 编译错误：
// u.name undefined (cannot refer to unexported field name)
// ============================================================
//
// 为什么会这样：
//   Go 的可见性由首字母大小写决定，不是关键字。
//   如果你想从其他包访问字段，字段名首字母必须大写。
//
// 如何修复：
//   type User struct {
//       Name string // 大写 = 导出
//       age  int    // 小写 = 私有（同一包内可访问）
//   }

package main

import "fmt"

// User 结构体演示可见性
type User struct {
	Name string // 大写 → 导出，其他包可见
	age  int    // 小写 → 私有，仅本包可见
}

// NewUser 构造函数（因为 age 私有，外部需要通过构造函数设置）
func NewUser(name string, age int) *User {
	return &User{Name: name, age: age}
}

// GetAge 通过导出方法访问私有字段
func (u *User) GetAge() int {
	return u.age
}

func main() {
	u := NewUser("Alice", 30)

	// 访问导出字段 → 正常
	fmt.Printf("Name: %s\n", u.Name)

	// 访问私有字段 → 正确（同一包内）
	fmt.Printf("Age (包内直接访问): %d\n", u.age)

	// 如果这段代码在其他包（如 otherpkg/），则：
	// fmt.Println(u.Name)    → ✅ 可以
	// fmt.Println(u.age)     → ❌ 编译错误！age 是小写
	// fmt.Println(u.GetAge()) → ✅ 通过导出方法访问
}
