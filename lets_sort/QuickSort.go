// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/6 - 6:22 下午

package main

import "fmt"

//快速排序
func main() {
	arr := []int{11,2,77,2,33,2,2,222,111}
	bubbleSort(&arr)
	fmt.Println(arr)
	BinarySearch(&arr,0,len(arr)-1,222)
}
func bubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
func BinarySearch(arr *[]int, leftIndex, rightIndex, value int) {
	if leftIndex > rightIndex {
		return
	}
	middle := (leftIndex + rightIndex) / 2 // 3 + 6 / 2
	//fmt.Println(middle)
	if (*arr)[middle] > value {
		// 如果中间值大于查找的值说明 我们需要查找的值还是在0到中间值之间所有我们还是继续查找
		// 但是是结束是中间值减一 因为中间值本身已经比较了
		BinarySearch(arr, leftIndex, middle-1, value)
	} else if (*arr)[middle] < value {
		// 如果中间值小于查找的值说明 我们需要查找的值还是不在0到中间值之间所有我们还是继续查找
		// 但是是起始位置是中间值加一 因为中间值本身已经比较了
		BinarySearch(arr, middle+1, rightIndex, value)
	} else {
		fmt.Println("找到了", value, "的下标为:", middle)
	}
}
