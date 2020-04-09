package main

import "fmt"

/*
	interface 是Go语言中提供的另一种数据类型
	它可以把所有共同性的方法定义在一起
	任何其他类型只要实现了这些方法就实现了这个接口
*/

//手机接口
type Phone interface {
	call() //约束别人需要实现的方法  方法签名
}

//不同品牌的手机 诺基亚
type Nokia struct {
}

//苹果
type iPhone struct{}

//实现接口的方法
func (n Nokia) call() {
	fmt.Println("诺基亚打电话了")
}
func (ip iPhone) call() {
	fmt.Println("苹果手机打电话了")
}
func main() {
	var phone Phone     //定义一个Phone类型的变量 interface
	phone = new(iPhone) //iPhone实现了这个接口所有就可以赋值
	phone.call()
	phone = new(Nokia)
	phone.call()
}
