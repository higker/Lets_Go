package main

import "fmt"

func main() {
	//map类型的切片
	var citys []map[string]string
	c1 := make(map[string]string, 1)
	c1["上海"] = "中国大陆"
	citys = append(citys, c1)
	fmt.Println(c1)
	c1 = make(map[string]string, 1)
	c1["纽约"] = "美国"
	citys = append(citys, c1)
	fmt.Println(citys[1])
	fmt.Println(citys)
	fmt.Printf("%[1]T %[1]v\n", citys)
	//map里面有切片
	m := make(map[string][]int, 10)
	m["one"] = []int{1, 2, 3, 4, 5, 5, 6}
	fmt.Printf("%[1]T %[1]v\n", m)
}
