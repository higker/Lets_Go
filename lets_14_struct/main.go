package main

import (
	"fmt"
)

//type 后面是类型名字
type MyInt int     //自定义类型
type yourInt = int //类型别名

//golang in struct
type User struct {
	name    string
	age     int8
	gender  rune
	address string
}

func main() {
	var n MyInt = 100
	var u yourInt = 100
	fmt.Println(u)
	fmt.Printf("%T %v\n", n, n)
	fmt.Println(User{"John", 21, '男', "上海"})
	var u1 User
	u1.name = "岳宁宁"
	u1.age = 25
	u1.address = "上海浦东新区"

	male := 0
	u1.gender = func() rune {
		if male == 0 {
			return '女'
		}
		return '男'

	}()
	fmt.Println(u1)
	//匿名结构体
	var s struct {
		name string
		age  int8
	}
	s.name = "Ding"
	s.age = 19
	fmt.Println("s = ", s)
}
