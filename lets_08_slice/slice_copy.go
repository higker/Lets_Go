package main

import "fmt"

//go语言slice中的copy函数的使用
func main() {
	/*
		由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。
		Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中
	*/
	s1 := make([]int, 5, 5)
	//s2 := s1
	//var s2 []int = s1 go语言会自己类型推断 所以不需要加类型声明
	var s2 = s1
	s2[3] = 100
	fmt.Printf("s1= %v \ns2= %v \n", s1, s2)
	/*
		这个为什么结果还是一样的?
		s1= [0 0 0 200 0]
		s2= [0 0 0 200 0]
		是因为上面我们通过等于把s2值等于s1的值了  在这个操作里面 s2 存储的地址和 s1值是一样的
		所以我们通过copy操作之后再去操作还是相当于操作的一个底层的数组
		除非重新建一个切片存储
	*/
	copy(s1, s2)
	s2[3] = 200
	fmt.Printf("s1= %v \ns2= %v \n", s1, s2)
	//var s3 []int //nil 这样声明的是空的 没有内存声明不进去
	s3 := make([]int, 5, 5) //这里已经在内存开辟新的空间
	copy(s3, s2)
	s3[3] = 300
	fmt.Printf("s3= %v \ns2= %v \n", s3, s2)
	t := []int{1, 2, 3, 4, 54, 5, 54}
	result, p := removeAsIndex(t, 0)
	fmt.Println(result, "-->remove value is-->", p)
	fmt.Println("cap(t)=", cap(t)) //切片是不存值的 值是由底层数组存在的 所以上面我们移除了指定的一个元素的值 只是移除了底层的数组的值 而不会影响数组的len 切片的原始的cap
	sINT := [...]int{1, 2, 3, 4, 5}
	fmt.Println(sINT)
	sl1 := sINT[:]
	fmt.Println(sl1)
	sl2 := make([]int, len(sl1), 10)
	fmt.Println(sl2)
	copy(sl2, sl1)
	fmt.Println(sl2)
}

func removeAsIndex(parm []int, index int) ([]int, int) {
	/*
		1. 拿到需要移除元素的下标
		2. parm[:lesIndex] 相当于 在内存里重新建了一个切片一样
			2.1 s := parm[:index] 这里就是从一个切片的第一个元素开始切 切到index下标的前一个元素为止
		3. parm[lesIndex+1:] 是从的指定的下标开始切 直到最后一个元素位置
			3.1  这里为什么要加一是因为我们不需要index这个元素 不包括
		4. append 的第一个参数是被追加源 第二个参数是追加源    可以理解为我们前面的学的slice切片 切除之后的切片
	*/
	value := parm[index]
	fmt.Printf("index type is: %T \n", index)
	parm = append(parm[:index], parm[index+1:]...)
	//return parm, parm[lesIndex] //1.这里的parm是截取元素之后的切片所以下面运行结果54
	return parm, value
}

/*
	面试结果 对应注释标签1
	[Running] go run "d:\Go_workspace\src\Lets_Go\lets_08_slice\slice_copy.go"
	s1= [0 0 0 100 0]
	s2= [0 0 0 100 0]
	s1= [0 0 0 200 0]
	s2= [0 0 0 200 0]
	s3= [0 0 0 300 0]
	s2= [0 0 0 200 0]
	lesIndex type is: int
	[1 2 3 54 5 54] 54
	[Done] exited with code=0 in 5.47 seconds
*/
