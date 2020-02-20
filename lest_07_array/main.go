/*
 * @Author: Deencode
 * @Date: 2020-01-07 19:55:58
 * @LastEditors  : BinScholl
 * @LastEditTime : 2020-01-07 20:30:56
 * @Description: go语言中的数组使用  Array
 * @Github: https://github.com/Deencode
 */

package main

import "fmt"

//数组的使用

var (
	people [3]string
)

func main() {
	//给people赋值
	for i := 0; i < len(people); i++ {
		people[i] = fmt.Sprintf("小米%d", i)
	}
	for _, v := range people {
		fmt.Println(v)
	}
	//声明一个5位元素位int64类型的数组并且初始化默认值
	ages := [5]int64{11, 16, 18, 19, 22}
	for i := 0; i < len(ages); i++ {
		if i == 0 {
			fmt.Print("ages = [\t")
		}
		//遍历打印ages
		fmt.Print(ages[i], "\t")
		if i == len(ages)-1 {
			fmt.Println("]")
		}
	}
	//比较方便的声明方式  ’...‘代表你不确定这个数组有几个元素位 让go语言自己去推断
	isGood := [...]bool{true, false, true}
	fmt.Println(isGood)
	//通过len函数来输出isGood数组里面有多少给元素
	fmt.Println(len(isGood))
	//没有初始化 数组会有默认值  前3个元素赋值了 后面2个没有赋值就有默认值
	nums := [5]int{1, 2, 3}
	fmt.Println(nums)
	status := [3]bool{true}
	fmt.Println(status)

	//多维数组 并且 数组类型要一至
	n := [3][2]int{
		[2]int{1, 2},
		[2]int{2, 2},
		[2]int{3, 2},
	}
	fmt.Println("方法二")
	//多维数组遍历 方式一
	for in, v1 := range n {
		fmt.Println("n[", in, "] = ", v1)
		for index, v2 := range v1 {
			fmt.Println("n[", in, "] [", index, "]= ", v2)
		}
	}
	fmt.Println("方法二")
	//多维数组遍历 方法二
	for i := 0; i < len(n); i++ { //先遍历第一层的元素
		fmt.Println("n[", i, "] = ", n[i])
		for u := 0; u < len(n[i]); u++ { //第一层元素里面是2个元素 所以继续遍历第二层的
			fmt.Println("n[", i, "] [", u, "]= ", n[i][u])
		}
	}
	fmt.Println(n)
	//访问n数组中的第二个元素里面的第二个元素
	fmt.Println(n[0][0]) //1
	fmt.Println(n[0][1]) //2
	fmt.Println()
	//通过下标去赋值 格式:下标:值
	floats := [10]float64{1: 1, 9: 9}
	fmt.Println(floats)
	//数组的长度是数组类型的一部分
	var array1 [3]int16
	var array2 [2]int16
	fmt.Printf("array1 : %T | array2 : %T\n", array1, array2)
	//fmt.Println(array1 == array2)数组类型不一样无法进行比较   因为在go语言中数组的长度是数组类型的一部分
	array3 := [2]int8{1, 2}
	array4 := [2]int8{1, 2}
	fmt.Println(array3 == array4)

	//数组是值类型的
	var b1 = [3]int8{1, 2, 3} //注意这里的等于号 ”=“ 如果在声明的时候就给数组赋值 要就加一个=
	var b2 [3]int8
	b2 = b1             //这里go语言会在内存开辟给空间 拷贝一份b1的值 给b2
	b2[1] = 100         //修改b2的下标为1的元素 值为100
	fmt.Println(b1, b2) //[1 2 3] [1 100 3]

}

/**
代码运行结果

[Running] go run "d:\Go_workspace\src\Lets_Go\lest_07_array\main.go"
小米0
小米1
小米2
ages = [	11	16	18	19	22	]
[true false true]
3
[1 2 3 0 0]
[true false false]
方法二
n[ 0 ] =  [1 2]
n[ 0 ] [ 0 ]=  1
n[ 0 ] [ 1 ]=  2
n[ 1 ] =  [2 2]
n[ 1 ] [ 0 ]=  2
n[ 1 ] [ 1 ]=  2
n[ 2 ] =  [3 2]
n[ 2 ] [ 0 ]=  3
n[ 2 ] [ 1 ]=  2
方法二
n[ 0 ] =  [1 2]
n[ 0 ] [ 0 ]=  1
n[ 0 ] [ 1 ]=  2
n[ 1 ] =  [2 2]
n[ 1 ] [ 0 ]=  2
n[ 1 ] [ 1 ]=  2
n[ 2 ] =  [3 2]
n[ 2 ] [ 0 ]=  3
n[ 2 ] [ 1 ]=  2
[[1 2] [2 2] [3 2]]
1
2

[0 1 0 0 0 0 0 0 0 9]
array1 : [3]int16 | array2 : [2]int16
true
[1 2 3] [1 100 3]

[Done] exited with code=0 in 3.292 seconds
*/
