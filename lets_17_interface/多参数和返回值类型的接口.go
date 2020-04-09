package main

import "fmt"

//接口定义 包含参数接口
type animal interface {
	move(int)
	eat()
}

type chicken struct {
	name string
}

func (c chicken) move(n int) {
	fmt.Println(c.name, "走了", n, "步.")
}
func (c chicken) eat() {
	fmt.Println(c.name, "吃东西.")
}

/*
	什么go语言的interface?
		1.interface是go语言中的数据类型
		2.它一些具有同样和共同性功能的函数定义到一起 并且还能起到约束作用
		3.其他如果想要实现接口就必须遵守interface里面的函数签名格式并且实现所有方法

*/
func main() {
	var aml animal
	aml = chicken{
		name: "小鸡",
	}
	aml.eat()
	aml.move(10)
}
