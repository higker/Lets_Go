package main

import "fmt"

type student struct {
	number string
	name   string
}

func main() {
	var s student
	s.name = "AAA"
	s.number = "123456"
	fmt.Println(s)
	fmt.Printf("s pointer=%p \n", &s)
	setName(s, "BBBB")
	fmt.Println(s)
	editName(&s, "CCC")
	fmt.Println(s)

	//初始化即赋值 并且获取指针
	var ss = &student{
		"123111",
		"DIng",
	}
	fmt.Printf("%T \n", ss)
	fmt.Println(&ss.name) //结构体里面的元素的地址都是连续一块的
	fmt.Println(&ss.number)
	sss := new(student)
	sss = &student{"12555", "Ubni"}
	fmt.Printf("%T %v \n", sss, *sss)
	fmt.Printf("%p %T\n", sss, sss)   //这里是sss存储的指针也就是我们创建student的内存地址
	fmt.Printf("%p %T\n", &sss, &sss) //sss自己本身的指针 指针自己也有自己的内存地址
	//fmt.Println(&sss == sss) //判断sss地址是否和sss本身存储的地址是否一样
}

//修改名字 这个修改只是内存的copy的副本没有修改本身
func setName(s student, n string) {
	fmt.Printf("s pointer=%p \n", &s)
	s.name = n
}

func editName(s *student, n string) {
	//语法糖  相当于 var d := *s   d.name = n
	s.name = n
}
