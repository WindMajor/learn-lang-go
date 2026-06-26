// WHAT: Bug 03 — 值接收者方法不能修改原始数据
// ERROR: 在方法上使用值接收者时，修改的是副本，不影响原始结构体
//        这是 Go 新手（尤其从 C++/Java 过来）最容易犯的错误
//
// ============================================================
// 运行结果：
// 修改前: {Alice 25}
// Bug 修改后: {Alice 25}  ← 没变！
// 正确修改后: {Bob 30}    ← 变了
// ============================================================
//
// 为什么会这样：
//   Go 中所有函数参数都是值传递，方法接收者也一样。
//   值接收者 `(u User)` 会拷贝整个结构体，修改的是拷贝。
//   只有指针接收者 `(u *User)` 才能修改原始数据。
//
// CONTRAST（与已知语言对比）：
//   - Rust: &self vs &mut self，编译器强制检查可变性
//   - TS:   class 的 this 总是指向原始对象（引用语义）
//   - Go:   值接收者是拷贝，指针接收者才是引用 —— 需要你手动选择！
//
//   规则：如果需要修改接收者，或者接收者很大（> 几十字节），
//         用指针接收者。小型不可变对象可用值接收者。
//
// 如何修复：
//   把 (u User) 改成 (u *User)

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// BUG: 值接收者 —— 修改的是副本
func (u User) BuggySetName(name string) {
	u.Name = name // 只修改了副本！
}

// 修复：指针接收者
func (u *User) FixedSetName(name string) {
	u.Name = name // 通过指针修改原始数据
}

func (u *User) FixedSetAge(age int) {
	u.Age = age
}

func main() {
	u := User{Name: "Alice", Age: 25}
	fmt.Printf("原始: %+v\n", u)

	u.BuggySetName("Bob") // 值接收者 → 修改无效！
	fmt.Printf("BuggySetName(\"Bob\") 后: %+v ← 没变！\n", u)

	u.FixedSetName("Bob") // 指针接收者 → 修改有效
	u.FixedSetAge(30)
	fmt.Printf("FixedSetName(\"Bob\") 后: %+v ← 变了！\n", u)
}
