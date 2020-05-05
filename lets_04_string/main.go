package main

import (
	"fmt"
	"strings"
)

const (
	s     = "string"
	c     = 'S'
	t     = 'T'
	r     = 'R'
	strMy = `
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
	fmt.Println(strMy)

	//字符串相关操作
	fmt.Println("字符串长度", len(strMy))

	//字符串拼接
	name := "Go"
	age := "20"
	ss := name + age
	fmt.Println(ss)
	//字符串格式化拼接 并且 格式化
	ss1 := fmt.Sprintf("我的名字叫%s 年龄%s", name, age)
	fmt.Println(ss1)
	//join可以将一个string切片拼接起来
	var ibyte []string = []string{"www", "ibyte", "com"}
	fmt.Println(strings.Join(ibyte, "."))
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
	//找出com最后一次出现的位置
	fmt.Println(strings.LastIndex(url, "com"))
	fmt.Printf("com最后出现的位置:%d", strings.LastIndex(url, "google"))
	tack := "www.tencent.com"
	for _, c := range tack {
		fmt.Printf("%c\n", c)
	}
	//同样达到单个输出字符的效果
	for i := 0; i < len(tack); i++ {
		fmt.Printf("%c\n", tack[i])
	}

	v3 := '中'
	fmt.Printf("%T,%d,%c,%q", v3, v3, v3, v3)
}
