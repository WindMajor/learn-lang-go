// Package calc 提供基础数学运算函数
// 本包用于演示 Go 的测试和工程化实践
package calc

import (
	"errors"
)

// Add 返回两数之和
func Add(a, b int) int {
	return a + b
}

// Subtract 返回 a - b
func Subtract(a, b int) int {
	return a - b
}

// Multiply 返回两数之积
func Multiply(a, b int) int {
	return a * b
}

// Divide 返回 a / b 的结果
// 如果 b 为 0，返回错误
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

// Max 返回两数中的较大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// IsPrime 判断 n 是否为质数（朴素算法，仅用于测试演示）
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fibonacci 返回第 n 个斐波那契数
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
