package main

import "fmt"

//slice 练习题
func main() {
	var s1 = make([]int, 3, 5)
	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
	}
	fmt.Println("s1 = ", s1) //0001234

	ss := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("ss=", ss)
	fmt.Println("ss len=", len(ss), "cap", cap(ss))
	ss = append(ss[:3], ss[4:]...)
	fmt.Println("ss=", ss)
	//cap=8 是应为切片存储数据是底层数组存储的 所以移除了一个元素 只是底层数组的对应值被后面一个元素覆盖了 底层数组的len是不变的
	fmt.Println("ss len=", len(ss), "cap", cap(ss))
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr = ", arr)
	sss := append(arr[:3], arr[4:]...)
	fmt.Println("sss=", sss)
	fmt.Println("arr len=", len(arr), "arr=", arr)
	fmt.Println("sss len=", len(sss), "cap", cap(sss))

	//练习题
	ss1 := []int{1, 2, 3}
	s2 := ss1
	var s3 []int //这里只是声明了一个变量 但是内存里面还没有开辟空间
	copy(s3, s1) //改成make 创建 make([]int,0,3)  但是这里的copy和append不一样 copy不能扩容
	fmt.Println(s2)
	s2[1] = 200
	fmt.Println("s2=", s2)
	fmt.Println("ss1=", ss1)
	fmt.Println("s3=", s3)

}

/*
[Running] go run "d:\Go_workspace\src\Lets_Go\lets_08_slice\slice_practice.go"
s1 =  [0 0 0 0 1 2 3 4]
ss= [1 2 3 4 5 6 7 8]
ss len= 8 cap 8
ss= [1 2 3 5 6 7 8]
ss len= 7 cap 8
arr =  [1 2 3 4 5]
sss= [1 2 3 5]
arr len= 5 arr= [1 2 3 5 5]
sss len= 4 cap 5
[1 2 3]
s2= [1 200 3]
ss1= [1 200 3]
s3= []
[Done] exited with code=0 in 3.092 seconds
*/
