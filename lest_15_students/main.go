package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//查看 新增 删除 学生
type student struct {
	id   int64
	name string
}

var (
	students = make([]student, 0, 100) //存放学生的slice
	mid      = -1                      //菜单id
)

func main() {
	menu()
}
func menu() {
	fmt.Println("=============学生系统============")
	fmt.Println("1.查看学生")
	fmt.Println("2.新增学生")
	fmt.Println("3.删除学生")
	fmt.Println("0.退出系统")
	fmt.Println("=================================")
	fmt.Println("请输入序号:")
	fmt.Scan(&mid)
	switch mid {
	case 1:
		see()
		menu()
	case 2:
		add()
		menu()
	case 3:
		delete()
		menu()
	case 0:
		exit()
		menu()
	default:
		menu()
	}
}

//查看学生
func see() {
	for n, v := range students {
		fmt.Println("--------------------------------------------------------------")
		fmt.Printf("第%d条记录| 学生ID: %d | 学生姓名: %s \n", n+1, v.id, v.name)
		fmt.Println("--------------------------------------------------------------")
	}
}

//添加学生
func add() {
	name := ""
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("请输入学生姓名:")
	fmt.Println("--------------------------------------------------------------")
	fmt.Scan(&name)
	s := NewStudent(name)
	students = append(students, s)
}

//删除学生
func delete() {
	flag := false
	id := int64(-1)
	fmt.Println("请输入要删除的学生ID:")
	fmt.Scan(&id)
	for i := 0; i < len(students); i++ {
		if id == students[i].id {
			//通过切片截取 移除掉不需要的那个学生信息
			students = append(students[:i], students[(i+1):]...)
			fmt.Println("--------------------------------------------------------------")
			fmt.Println("删除学生成功！")
			fmt.Println("--------------------------------------------------------------")
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("没有这个学生信息!!")
		fmt.Println("--------------------------------------------------------------")
	}

}

//构造函数
func NewStudent(name string) student {
	rand.NewSource(time.Now().UnixNano())
	id := rand.Int63()
	return student{
		id:   id,
		name: name,
	}
}

//退出程序
func exit() {
	os.Exit(0) //正常运行程序并退出程序；
}
