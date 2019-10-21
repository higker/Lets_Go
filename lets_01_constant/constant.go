package main

import "fmt"

//常量的使用 和 定义
//常量在程序运行期间是不能改变变量值的！！！
const (
	pi = 3.1415927
)

//如果批量声明常量时,如果某一行没有赋值,默认值就和上一行一致
const (
	t1 = 1000
	t2
	t3
)

//go 语言常量的使用
func main() {
	//pi = 4 这里就直接报错了 运行期间不能改成常量的值
	fmt.Println(pi)
	fmt.Println(t1, t2, t3)

}
