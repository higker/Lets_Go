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
	var ptrT *int
	ptrT = &i //这里我们把i的地址给了ptrT这个指针了所以运行是没有问题 如果是上面那种直接操作 是nil的~~
	*ptrT = 999
	fmt.Println(*ptrT)
	fmt.Println(ptrT)

	//也可以这样写

	var ptr2 *int = new(int) //new 关键字不是初始化对象 而是在帮助引用类型在内存里面开辟空间并且赋零值
	//ptr2 = 999
	fmt.Println(*ptr2)
}
