package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//学生结构体
type student struct {
	SId   int    `json:"s_id"`
	SName string `json:"s_name"`
}

//实例化学生构造函数
func newStudent(id int, name string) student {
	return student{
		SId:   id,
		SName: name,
	}
}

//学生管理者结构体
type studentsManage struct {
	allStud map[int]student
}

//初始化管理者
func initMgr() *studentsManage {
	return &studentsManage{
		//初始化数据库
		allStud: make(map[int]student, 100),
	}
}

//添加用户
func (sm studentsManage) add() {
	rand.NewSource(time.Now().UnixNano())
	var name string
	fmt.Println("请输入学生姓名:")
	fmt.Scan(&name)
	id := rand.Int()
	sm.allStud[id] = newStudent(id, name) //存入到模拟的map数据库里
}

//查看用户
func (sm studentsManage) see() {
	for _, v := range sm.allStud {
		fmt.Println("--------------------------------------------------------------")
		fmt.Printf("| 学生ID: %d | 学生姓名: %s \n", v.SId, v.SName)
		fmt.Println("--------------------------------------------------------------")
	}
}

//删除用户
func (sm studentsManage) del() {
	var sid int
	fmt.Println("请输入学生SID:")
	fmt.Scan(&sid)
	_, isOK := sm.allStud[sid] //判断用户是否存在
	if !isOK {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("没有这个学生信息!!")
		fmt.Println("--------------------------------------------------------------")
	} else {
		delete(sm.allStud, sid) //通过内置del函数删除用户
		students = append(students[:i], students[(i+1):]...)
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("删除学生成功！")
		fmt.Println("--------------------------------------------------------------")
	}
}

//退出系统
func (sm studentsManage) exit() {
	os.Exit(0) //正常运行程序并退出程序；
}

func (sm studentsManage) menu() {
	mid := -1 //菜单id
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
		sm.see()
		sm.menu()
	case 2:
		sm.add()
		sm.menu()
	case 3:
		sm.del()
		sm.menu()
	case 0:
		sm.exit()
		sm.menu()
	default:
		sm.menu()
	}
}
func main() {
	//初始化
	sm := initMgr()
	sm.menu()
}
