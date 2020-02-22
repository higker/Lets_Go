package main

import "fmt"

//通过make函数创建slice
/*
	1. go语言切片不保存数据 数据全部是由底层数组保存   切片就是一个框 你要那块 就去切数组的那块
	2. 切片的len 是切片的元素个数   切片cap的值是根据 需要截取的起点index到原始数组最后一个元素的位数
*/
func main() {
	s1 := make([]int, 6, 10)
	fmt.Printf("s1= %v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := make([]string, 5) //这里make函数默认需要传入3个参数的   make(切片类型,切片长度,切片容量)
	fmt.Printf("s2= %v,len(s2)= %d,cap(s2)= %d\n", s2, len(s2), cap(s2))
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s3 := a1[3:8]
	fmt.Println("a1 = ", a1)
	fmt.Println("s3 = ", s3, "len(s3) = ", len(s3), "cap(s3) = ", cap(s3))
	fmt.Println("a1 == nil ?", len(a1) == 0) //判断一个切片是否为空 得使用len函数来判断
	//fmt.Println("a1 == nil ?", a1 == nil) 比较切片是否为空 不能使用 == nil
	ss := a1[:]
	subSliceOf(ss)
	subSlice(ss)

	//切片的赋值
	var name = [2]string{"马化腾", "扎克伯格"} //名字数组
	var person = make([]string, 10)     //人名切片
	//此时的内存引用已经丢弃了原来通过make函数创建的person底层那个数组
	//现在而是通过我们刚才自己定义name这个数组来进行操作
	//person = name[:1] person: [马化腾] len(person)= 1 cap(person)= 2
	person = name[:]
	fmt.Println("person:", person, "len(person)=", len(person), "cap(person)=", cap(person))

	var city = []string{"上海", "深圳", "北京", "天津", "重庆"}
	fmt.Println("city =", city)
	city = city[1:4]
	fmt.Println("city = ", city, "city的len是", len(city), "city的cap是", cap(city))
}

/*
	go语言中的slice遍历
*/

func subSliceOf(s []int) {
	for i := 0; i < len(s); i++ {
		fmt.Println("index:", i, ",value:", s[i])
	}
}
func subSlice(s []int) {
	for i, v := range s {
		fmt.Println("index:", i, ",value:", v)
	}
}
