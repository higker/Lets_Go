// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/15 - 4:21 下午

package lets_36_testing

import (
	"testing"
)

func TestMath(t *testing.T) {
	// 自定义测试结构体
	type MathCase struct {
		n, m, result int
	}
	// 自定义子测试map
	testGroup := map[string]MathCase{
		"add":   {1, 2, 3},
		"sub":   {3, 1, 2},
		"multi": {3, 2, 6},
		"div":   {6, 2, 3},
	}
	// 测试执行函数
	for name, mathCase := range testGroup {
		t.Run(name, func(t *testing.T) {
			s := -1
			switch name {
			case "add":
				s = Add(mathCase.n, mathCase.m)
			case "sub":
				s = Sub(mathCase.n, mathCase.m)
			case "multi":
				s = Multi(mathCase.n, mathCase.m)
			case "div":
				s = Div(mathCase.n, mathCase.m)
			default:
				t.Fatalf("No executable testing name :%s", name)
				return
			}
			if mathCase.result != s {
				t.Fatalf(" add computer result error， want %d , got %d", mathCase.result, s)
			}
		})
	}
}
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, b.N)
		//fmt.Println(add)
	}
}

// fib.go

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// fib_test.go

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
