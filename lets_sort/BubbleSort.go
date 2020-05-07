// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/6 - 6:19 下午

package main

import "fmt"

type Array []int

//冒泡排序 bubble sort
func main() {
	var array = [6]int{67, 32, 88, 12, 5, 46}
	var arr = []int{67, 32, 88, 12, 5, 46, 1024, 223, 552, 15, 99, 367, 208}
	BubbleSort(&array)
	BubbleSortOfSlice(&arr)
	fmt.Println(array)
	fmt.Println(arr)
	var ars = Array{67, 32, 88, 12, 5, 3, 46, 1024, 223, 552, 15, 99, 367, 208}
	ars.ofBubbleSort()
	fmt.Println(ars)
	sort := symbolBubbleSort('<')
	var arr2 = []int{7, 32, 8, 12, 5, 46, 1024, 123, 552, 15, 99, 367, 208}
	sort(&arr2)
	fmt.Println(arr2)
}

func BubbleSort(arr *[6]int) {
	temp := 0
	for i := 0; i < 5; i++ {
		temp = arr[i]
		if arr[i] > arr[i+1] { // 如果前面一个数大于后面一个数这个条件就成立
			arr[i] = arr[i+1] // 条件成立就把他们的位置做交换 把大的那个数往后面移动一位
			arr[i+1] = temp   // 数字小的就往前移动一位
		}
	}
}

//Invalid operation: 'arr[i]' (type '*[]int' does not support indexing
func BubbleSortOfSlice(arr *[]int) {
	for j := len(*arr) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if (*arr)[i] > (*arr)[i+1] { // 如果前面一个数大于后面一个数这个条件就成立
				(*arr)[i], (*arr)[i+1] = (*arr)[i+1], (*arr)[i]
			}
		}
	}
}

func symbolBubbleSort(r rune) func(arr *[]int) {

	return func(arr *[]int) {
		for i := 0; i < len(*arr); i++ {
			for j := 0; j < len(*arr)-1-i; j++ { //长度减1是因为防止下标越界 减i是每次已经排序好了几位就不用排序了
				switch r {
				case '>':
					if (*arr)[j] > (*arr)[j+1] {
						(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
					}
				default:
					if (*arr)[j] < (*arr)[j+1] {
						(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
					}
				}
			}
		}
	}

}

func (arr *Array) ofBubbleSort() {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1-i; j++ { //长度减1是因为防止下标越界 减i是每次已经排序好了几位就不用排序了
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
