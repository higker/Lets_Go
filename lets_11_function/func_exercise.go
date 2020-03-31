package main

import (
	"fmt"
	"strings"
)

//闭包操作练习题
func main() {
	var checkGIF func(string) string = makeEnding(".gif")
	fmt.Println(checkGIF("哈哈哈"))
	checkJPG := makeEnding(".jpg")
	fmt.Println(checkJPG("呵呵.jpg"))

	add, sub := buildMath(100)
	fmt.Println(add(10)) //110
	fmt.Println(sub(10)) //90 是因为我们这里是直接返回的 没有共用一个
}

/*
	高阶函数 & 闭包操作
	第一步操作是 通过传入文件后缀名 来构建一个 检查文件是否相关后缀的函数
	执行顺序: 1. 传入 suffix  2.会缓存在内存里 3.返回一个函数用来检查是否对应后缀名
*/
func makeEnding(suffix string) func(file string) string {
	return func(file string) string {
		if strings.HasSuffix(file, suffix) { //检查file是否以suffix为后缀
			return file
		}
		return file + suffix
	}
}

func buildMath(number int) (add func(n int) int, sub func(nu int) int) {
	/*
		//加法实现
		add = func(n int)int{
			return numder+n
		}
		//减法实现
		sub = func(nn int)int{
			return numder-nu
		}
	*/
	return func(n int) int {
			return number + n
		}, func(nu int) int {
			return number - nu
		}
}
