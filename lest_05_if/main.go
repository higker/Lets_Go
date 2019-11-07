package main

import "fmt"

var (
	age    = 18
	result = 0
)

//go 里面的 条件 语句 if example
func main() {

	if age > 18 {
		fmt.Println("你可以找女朋友了")
	} else {
		fmt.Println("你还小呢")
	}
	fmt.Println(compare(2000, 2000+1))
}
func compare(num1 int64, num2 int64) (result int64) {
	if num1 > num2 {
		result = num1
	} else if num2 > num1 {
		result = num2
	} else {
		fmt.Println("相等")
	}
	return result
}
