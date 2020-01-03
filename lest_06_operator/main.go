/*
 * @Author: BinScholl
 * @Date: 2019-12-13 20:38:08
 * @LastEditors: BinScholl
 * @LastEditTime: 2019-12-21 19:51:15
 * @Description: go语言中的运算符 加减乘除法 取余
 * @Github: https://github.com/BinScholl
 */

package main

import "fmt"

//go 语言运算符
func main() {
	/*
	 * 算数运算符:
	 * 加+ 减-  乘* 除/(除法取商)  取余% 自加加++ 自减减--
	 */
	add1 := 200
	add2 := 101
	fmt.Printf("%d + %d = %d\n", add1, add2, (add1 + add2))

	sub1 := 301
	sub2 := 200
	fmt.Printf("%d - %d = %d\n", sub1, sub2, (sub1 + sub2))

	multi1 := 100
	multi2 := 2
	fmt.Printf("%d * %d = %d\n", multi1, multi2, (multi1 * multi2))

	div1 := 100
	div2 := 2
	fmt.Printf("%d / %d = %d\n", div1, div2, (div1 / div2))

	div3 := 100
	div4 := 3
	//%代表取余     2个百分号 : ——>是因为%格式化缘故取余
	fmt.Printf("%d %% %d = %d\n", div3, div4, (div3 % div4))

	fmt.Println(add(0))
	fmt.Println(sub(1))
}

/**
 * @description: 一个数加加
 * @param int64 number
 * @return: number add sum
 */
func add(num int64) (sum int64) {
	num++
	return num
}

//一个数减减
func sub(num int64) (result int64) {
	num--
	return num
}

// [Running] go run "/Users/ding/Documents/GO_CODE_DEV/src/Lets_Go/lest_06_operator/main.go"
// 200 + 101 = 301
// 301 - 200 = 501
// 100 * 2 = 200
// 100 / 2 = 50
// 100 % 3 = 1
// 1
// 0

// [Done] exited with code=0 in 1.514 seconds
