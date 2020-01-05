/*
 * @Author: BinScholl
 * @Date: 2020-01-03 19:58:26
 * @LastEditors  : BinScholl
 * @LastEditTime : 2020-01-05 13:52:12
 * @Description: go语言中goto关键字的使用
 * @Github: https://github.com/BinScholl
 */

package main

import "fmt"

func main() {
	gotoDemo3()
}

//通常写法1
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'D' {
				// 设置退出标签
				breakFlag = true
				break //跳出内层循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		// 外层for循环判断  跳出外层循环
		if breakFlag {
			break
		}
	}
	fmt.Println("我是goto1我还会执行吗?")
}

//尝试使用return语句实现
//！！！这种方式可以实现跳出 for 但是后面的代码将无法进行执行
func gotoDemo3() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'D' {
				// 设置退出标签
				return
			}
			fmt.Printf("%v-%c\n", i, j)
		}

	}
	fmt.Println("我是goto3我还会执行吗?")
}

//通过goto实现
//goto可以指定程序执行到某个条件时 直接跳转到goto标签锁指定的代码
func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 9 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return

breakTag: // 标签名称 这个就是定义的label name 方便goto找到这个标签代码块去执行!
	fmt.Println("for循环都跳出")

	fmt.Println("我是goto3我还会执行吗?")
}
