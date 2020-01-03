/*
 * @Author: BinScholl
 * @Date: 2020-01-03 15:05:14
 * @LastEditors  : BinScholl
 * @LastEditTime : 2020-01-03 15:26:51
 * @Description go语言使用for循环打印九九乘法表 MultiplicationTable
 * @Github: https://github.com/BinScholl
 */

package main

import "fmt"

func main() {
	// 遍历, 决定处理第几行
	for i := 1; i < 10; i++ {
		// 遍历, 决定这一行有多少列
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		//手动回车
		fmt.Println()
	}
	fmt.Println("倒序 九九乘法表")
	for i := 9; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
