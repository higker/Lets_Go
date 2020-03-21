package main

import "fmt"

func main() {
	n := "100"
	// var ptr *int
	i := 200
	ptr := &i
	fmt.Println("n 的pointer", ptr)
	fmt.Println("i=", i)
	n = fmt.Sprintf("%v", *ptr)
	fmt.Println("n=", n)
	//fmt.Println(&n == &i)

	//var ptrTemp *int 这里声明了一个int指针类型的变量   *代表的是取指针的值 但是这个值nil 没有内存地址
	//*ptrTemp = 100 所以到这里来通过*取内存地址的值就会报错
	//fmt.Println(ptrTemp)
}
