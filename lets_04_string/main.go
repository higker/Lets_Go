package main

import (
	"fmt"
	"strings"
)

const (
	s      = "string"
	c      = 'S'
	t      = 'T'
	r      = 'R'
	string = `
		My name is "Ding Shuo",
		My age is 20
		Useing langues Chinese
  	`
)

func main() {
	fmt.Println(s)
	fmt.Println(c)
	fmt.Println(t)
	fmt.Println(r)
	fmt.Println(string)

	//字符串相关操作
	fmt.Println("字符串长度", len(string))

	//字符串拼接
	name := "Go"
	age := "20"
	ss := name + age
	fmt.Println(ss)
	//字符串格式化拼接 并且 格式化
	ss1 := fmt.Sprintf("我的名字叫%s 年龄%s", name, age)
	fmt.Println(ss1)

	url := "www.google.com"
	result := strings.Split(url, ".")
	fmt.Println(result)
	//是否包含某个字符串
	var isContains bool = strings.Contains(url, "google")
	fmt.Println(isContains)
	//是否以什么开头 前缀
	fmt.Println(strings.HasPrefix(url, "www"))
	//是否以什么结尾 后缀
	fmt.Println(strings.HasSuffix(url, "com"))
	//判断子串出现的位置
	fmt.Println(strings.Index(url, "google"))

}
