package main

import "fmt"

func main() {
	/*
			go语言中的切片 slice
			1. 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
			2. 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
			3. go语言中的slice 是引用类型
			值类型：保存在线程栈上的，由系统自动释放资源

		引用类型：保存在托管堆中的，引用类型包括类、接口、委托和装箱值类型reference type

		引用规则：

		（1）引用被创建的同时必须被初始化（指针则可以在任何时候被初始化）。

		（2）不能有NULL引用，引用必须与合法的存储单元关联（指针则可以是NULL）。

		（3）一旦引用被初始化，就不能改变引用的关系（指针则可以随时改变所指的对象）。

		作者：派大星_5274
		链接：https://www.jianshu.com/p/d0f615069d96
		来源：简书
		著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	*/

	//声明一个slice 定义了一个int类型的切片
	var s1 []int
	var s2 []string
	fmt.Println("s1:", s1, "s2:", s2)
	//nil在go语言中就是代表空的意思 在内存里面还没有开辟空间
	fmt.Println("s1 == nil ?", s1 == nil)
	//初始化
	s1 = []int{1, 2, 3, 4, 5, 6}
	s2 = []string{"上海", "深圳", "武汉"}
	fmt.Println("s1:", s1, "s2:", s2)
	//这里的s1已经在内存里面开辟了空间 所以这里不是空
	fmt.Println("s1 == nil ?", s1 == nil)

	/*

		2.切片的长度和容量
		切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	*/
	fmt.Printf("len:%d;cap:%d\n", len(s1), cap(s1))

	//由一个已经存在的数组来得到slice
	arr := [...]int{1, 23, 2, 34, 6}
	//这里就是从arr数组里面通过下标来截取或者切取 从index为0的元素开始切取3给元素 然后赋值给arr2
	arr2 := arr[0:3]
	fmt.Println("arr2:", arr2)
	arr3 := arr[1:len(arr)]
	fmt.Println("arr3:", arr3)
	s3 := arr[:4]  //从第一个切割到第4
	s5 := arr[1:3] //从第1个切割到4
	s6 := arr[3:]  //从第3个切割到最后
	s7 := arr[:]   //从第一个切割到最后
	fmt.Println(s3, s6, s7)
	//切片的容量是值得是底层数组的容量   左闭右包  从指定的下标到原始长度或者容量
	/*
		帮助链接:https://www.liwenzhou.com/posts/Go/06_slice/
	*/
	fmt.Printf("len(s3):%d;cap(s3):%d    len(arr):%d\n", len(s3), cap(s3), len(arr))
	fmt.Println(s3)
	fmt.Println(s5)
	fmt.Printf("len(s5):%d;cap(s5):%d\n", len(s5), cap(s5))
	fmt.Println("arr:", arr)
	arr[4] = 66
	fmt.Println("arr:", arr)
	s8 := arr[3:]                                                    // [34 66]
	fmt.Println("s8:", s8, "len(s8)=", len(s8), "cap(s8)=", cap(s8)) // len=2 cap=2
}

/*
	以上代码运行结果:

[Running] go run "d:\Go_workspace\src\Lets_Go\lets_08_slice\slice.go"
s1: [] s2: []
s1 == nil ? true
s1: [1 2 3 4 5 6] s2: [上海 深圳 武汉]
s1 == nil ? false
len:6;cap:6
arr2: [1 23 2]
arr3: [23 2 34 6]
[1 23 2 34] [34 6] [1 23 2 34 6]
len(s3):4;cap(s3):5    len(arr):5
[1 23 2 34]
[23 2]
len(s5):2;cap(s5):4
arr: [1 23 2 34 6]
arr: [1 23 2 34 66]
s8: [34 66] len(s8)= 2 cap(s8)= 2

[Done] exited with code=0 in 4.621 seconds

*/
