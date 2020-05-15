// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/15 - 4:21 下午

package lets_36_testing

import "testing"

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
