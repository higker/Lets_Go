package main

import (
	"fmt"
	"reflect"
)

//批量声明变量并且指定类型
//go语言推荐使用驼峰命名 student_name  StudenName
var (
	name        string
	age         int
	isMan       bool
	studentName string
)

//go语言变量实例
func main() {
	// 通过var 关键字 go 自动类型推断 变量声明并且赋值
	var varNum = 1111
	var varStr = "String"
	var varBool = false
	fmt.Println(varNum)
	fmt.Println(varStr)
	fmt.Println(varBool)

	//给var代码块的变量赋值
	name = "Go language"
	//age = "2009" 指定类型 强行赋值其他类型会报错
	age = 2009
	isMan = true
	//判断studentName 的值是否为空
	if reflect.ValueOf(studentName).IsValid() {
		studentName = "Ding Shuo"
	}
	fmt.Println(name, age, isMan, studentName)
	//这个相当于 var time = "2019年10月20日21:42:15""
	time := "2019年10月20日21:41:54"
	fmt.Println(time)

	//variable := "variable" 这种只能在函数里使用  简单声明

	//不接收第二个返回的变量 _代表匿名变量
	ayo, _ := test(2019)
	fmt.Println(ayo)
	//不接收第一个 不占用内存
	_, ayo2 := test(2020)
	fmt.Println(ayo2)
}

//test测试匿名变量 参数   多路返回值
func test(num int) (int, string) {
	return num, "匿名变量测试"
}
