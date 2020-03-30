package main

import "fmt"

func main() {
	func1()
	func2("Jesse Deen")
	str := func3()
	fmt.Println("str=", str)
	func4("DING", 1, 2, 3, 4)
	sum := sum(1, 2)
	fmt.Println(sum)
	testDefer()
	testFunc(x)       //去掉括号
	fun := buildSum() //Go里面函数式编程
	fmt.Printf("fun typs = %T \n", fun)
	s := fun(1, 3) //这个fun 是通过buildSum函数返回的 返回值是一个sum函数类型
	fmt.Println(s)
}

//高阶函数
//函数类型作为参数
func testFunc(x func() int) { //x 是一个有int返回值的函数类型
	result := x()
	fmt.Println(result)
}

func x() int {
	return 6
}

// 返回值是函数类型的
func buildSum() func(a, b int) (s int) {
	s := sum
	return s
}

//defer 函数
func testDefer() {
	fmt.Println("1.打开数据库连接")
	defer fmt.Println("3.关闭数据库连接")
	fmt.Println("2.操作数据库")
}

//go语言定义一个无参 无返回值的函数
func func1() {
	fmt.Println("func1()")
}

//有参数
func func2(name string) {
	fmt.Println("Hello ", name)
}

//有返回值return
func func3() string {
	return "Jesse Deen"
}

//有命名返回值的
func sum(i, n int) (sum int) { //这里的sum 就是代表为sum 已经声明好返回类型和返回类型的参数名字
	sum = i + n
	return
}

//可变类型
func func4(name string, y ...int) {
	//y就一个切片
	fmt.Println("name:", name, "Y=", y)
	for _, v := range y {
		fmt.Println(v)
	}
}
