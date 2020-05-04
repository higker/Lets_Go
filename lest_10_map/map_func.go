package main

import (
	"fmt"
)

type book struct {
	name  string
	about string
}

//go语言里面的map指类型也可以是function类型
func main() {

	//基本数据类型和引用类型map例子
	strMap := make(map[string]int, 100)
	strMap["计算机科学"] = 100
	strMap["英文"] = 90
	strMap["数学"] = 119
	fmt.Println(strMap)
	//引用类型 1001类型就代表我国4大名著
	book1 := book{"西游记", "吴承恩"}
	book2 := book{"三国演义", "罗贯中"}
	books := []book{book1, book2}
	BookMap := make(map[int][]book, 10)
	BookMap[1001] = books //赋值操作
	fmt.Println(BookMap)
	fmt.Println(BookMap[1001][1]) //访问BookMap里面的元素

	//value是function类型的Map
	funcMap := make(map[string]func(x, y int) int, 100)
	funcMap["add"] = func(x, y int) int {
		return x + y
	}
	funcMap["sub"] = func(x, y int) int {
		return x - y
	}
	funcMap["div"] = func(x, y int) int {
		return x / y
	}
	funcMap["mult"] = func(x, y int) int {
		return x * y
	}
	add := funcMap["add"]
	fmt.Println(add(1, 2))
	fmt.Println(funcMap["sub"](4, 2))
	fmt.Println(funcMap["mult"](2, 8))
}
