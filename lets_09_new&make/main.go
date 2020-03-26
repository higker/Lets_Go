package main

import (
	"fmt"
)

//go语言中的内置new函数和make函数的使用
func main() {
	var ptr1 *int //声明一个int类型的指针
	//*ptr1 = 100        //获取到ptr1指针地址的值并且赋值为100
	//fmt.Println(*ptr1) //运行报错 因为上的var 声明的但是未在内存里面开辟地址
	//*ptr1 = new(int) //这样取也是取不到的应为*ptr还是空的里面没有值  取不到值
	ptr1 = new(int) //通过new函数可以初始化并且开辟内存空间 默认值为对应类型的零值
	*ptr1 = 100
	fmt.Println(*ptr1)

	/*
		make 和 new的区别
		new : 是用于基本数据类型的 给你基本数据类型申请内存空间 并且返回相应类型的指针 (string int float bool)
		make : 是用于 map , slice , chan申请内存空间的 并且返回相应的数据类型本身
	*/

	//new 使用实例
	var str *string = new(string)
	fmt.Println("str= ", *str)
	*str = "Hello" //取str指针中的值
	fmt.Println(*str)
	ofStr := *str
	fmt.Printf("ofStr = %s , %T\n", ofStr, ofStr)

	//make 使用实例
	//错误使用
	var sli []int = []int{}  //这样声明的切片是空的
	fmt.Println("sli=", sli) // sli = nil
	fmt.Println(sli == nil)
	fmt.Println(len(sli))
	fmt.Println("cap(sli)=", cap(sli), "len(sli)=", len(sli))
	//正确使用
	var sliTow = make([]int, 5, 10)
	fmt.Println("sliTow=", sliTow)
	//声明一个int类型的切片指针 用来存储int类型的指针
	var ptrSli = new(*int)
	*ptrSli = &sliTow[1]
	fmt.Println(ptrSli)
	//有趣的一道题
	//https://studygolang.com/articles/13595?fr=sidebar
}
