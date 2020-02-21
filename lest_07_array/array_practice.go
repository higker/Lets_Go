package main

import "fmt"

func main() {
	//作业 完成
	//1.计算一个数组里面所以元素的和
	var testArr = [3]int{1, 2, 3}
	fmt.Println(sumArray(testArr))

	//2.找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	var testArr2 = [5]int{2, 3, 4, 1, 2}
	var testArr3 = [5]int{1, 3, 1, 1, 2}
	subscript(testArr2, 5)
	subscript(testArr3, 2)
	subscriptOf(testArr3, 2)
}

//计算一个数组里面所有元素的和 PS:这里数组元素位一定要和被传入的数组类型匹配 因为go语言中的数组的元素位是数组类型的一部分
/**
  这样写是错误的
  func sumArray(arr [...]int) int {
	  ...
  }
  必须指定数组类型 切记在go语言中数组长度是数组类型的一部分
*/
func sumArray(arr [3]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

//找出2个元素和相加为n的 元素下标
func subscript(arr [5]int, sum int) {
	for i, v := range arr {
		for n, val := range arr {
			if i == n {
				break
			}
			if v+val == sum {
				fmt.Println("元素相加和为", sum, "的元素下标有:{index:", i, "value:", v, ",index:", n, "value:", val, "}")
			}
		}
	}
}

//找出2个元素和相加为n的 元素下标
func subscriptOf(arr [5]int, sum int) {
	fmt.Println("普通循环实现2:")
	//外循环从第一个元素开始
	for i := 0; i < len(arr); i++ {
		//内循环从外循环的后面一个开始  这样是因为防止自己和自己比较
		for n := i + 1; n < len(arr); n++ {
			if (arr[i] + arr[n]) == sum {
				fmt.Println("元素相加和为", sum, "的元素下标有:{index:", i, "value:", arr[i], ",index:", n, "value:", arr[n], "}")
			}
		}
	}
}
