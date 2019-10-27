package main

import "fmt"

//fmt占位符
func main() {
	var n int = 100
	fmt.Printf("查看类型%T\n", n)
	fmt.Printf("查看值%v\n", n)
	fmt.Printf("查看值并且带上类型描述%#v\n", n)
	fmt.Printf("二进制%b \n", n)
	fmt.Printf("十进制%d\n", n)
	fmt.Printf("八进制%o\n", n)
	fmt.Printf("十六进制%x\n", n)
	fmt.Printf("字符串%s\n", "字符串")
}
