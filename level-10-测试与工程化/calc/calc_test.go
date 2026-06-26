// WHAT: calc 包的测试文件
// WHY: 展示 Go 的表驱动测试、子测试、基准测试
// 文件名以 _test.go 结尾，Go 工具链自动识别

package calc

import (
	"fmt"
	"testing"
)

// TestAdd 表驱动测试 — Go 社区的惯用测试范式
func TestAdd(t *testing.T) {
	// 定义测试用例表
	tests := []struct {
		name string // 用例名称（用于 t.Run 和报告）
		a, b int
		want int
	}{
		{"两个正数", 1, 2, 3},
		{"正数加零", 5, 0, 5},
		{"负数加正数", -1, 1, 0},
		{"两个负数", -3, -7, -10},
		{"大数相加", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		// t.Run 创建子测试 — 每个用例独立运行
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestDivide 除法测试（含错误案例）
func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"正常除法", 10, 2, 5, false},
		{"小数除法", 7, 3, 7.0 / 3.0, false},
		{"除数为零", 10, 0, 0, true},
		{"零除以任何数", 0, 5, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Divide(%f, %f) 期望错误但没有", tt.a, tt.b)
				}
				return
			}
			if err != nil {
				t.Errorf("Divide(%f, %f) 不期望错误: %v", tt.a, tt.b, err)
				return
			}
			if got != tt.want {
				t.Errorf("Divide(%f, %f) = %f, want %f", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestIsPrime 质数判断测试
func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{17, true},
		{97, true},
		{100, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			if got := IsPrime(tt.n); got != tt.want {
				t.Errorf("IsPrime(%d) = %t, want %t", tt.n, got, tt.want)
			}
		})
	}
}

// TestFibonacci 斐波那契测试
func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, want := range expected {
		got := Fibonacci(i)
		if got != want {
			t.Errorf("Fibonacci(%d) = %d, want %d", i, got, want)
		}
	}
}

// BenchmarkAdd 基准测试 — 函数名以 Benchmark 开头
func BenchmarkAdd(b *testing.B) {
	// b.N 由测试框架自动调整，直到结果稳定
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

// BenchmarkFibonacci 斐波那契基准测试
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

// BenchmarkIsPrime 质数判断基准测试
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(999983) // 一个大质数
	}
}
