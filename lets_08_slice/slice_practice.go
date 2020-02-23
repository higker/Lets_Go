package main

import "fmt"

//slice 练习题
func main() {
	var s1 = make([]int, 3, 5)
	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
	}
	fmt.Println("s1 = ", s1) //0001234
}
