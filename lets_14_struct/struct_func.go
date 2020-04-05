package main

import "fmt"

type auto struct {
	name string
	seat uint8
}

func main() {
	a := NewAuto("lynk&co", 5)
	fmt.Println(a.name, a.seat)
	a.Start()
	//接收者同时用表示调用该方法具体用那个类型
	NewAuto("法拉利488", 2).Start()
}

// 创建汽车 返回汽车的内存地址也就是一个*auto的指针
func NewAuto(name string, seat uint8) *auto {
	return &auto{
		name: name,
		seat: seat,
	}
}

//接收者 a就指定谁能调用这个方法  类似于Java中点出函数来使用的效果
func (a auto) Start() {
	fmt.Println(a.name + "点火启动了~")
}
