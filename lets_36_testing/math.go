// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/15 - 4:16 下午

package lets_36_testing

func init() {
	Math := make(map[string]func(n, m int) int, 4)
	Math["add"] = Add
	Math["sub"] = Sub
	Math["multi"] = Multi
	Math["div"] = Div
}

func Add(n, m int) int {
	return n + m
}
func Sub(n, m int) int {
	return n - m
}
func Multi(n, m int) int {
	return n * m
}
func Div(n, m int) int {
	return n / m
}
