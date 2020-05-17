// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/17 - 11:14 上午

// 使用go操作数据库 增删改查
package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"strings"
)

var (
	dbPool     *sql.DB //数据库连接池结构体
	MaxConnNum = 10    // 连接池最大连接数
	funcNum    = "0"   //功能序号
	reader     = bufio.NewReader(os.Stdin)
	es         = &employee{"", "", -1}
)

type DbConfig struct {
	user, pwd, addr, database string
	port                      int
}

type employee struct {
	name, password string
	id             int
}

func (e *employee) insert(em *employee) {
	exec, err := dbPool.Exec("INSERT INTO users (name,password) VALUES (?,?)", em.name, em.password)
	if err != nil {
		fmt.Println("==========================================================")
		fmt.Println("add employee info failed.")
		fmt.Println("==========================================================")
	}
	if num, _ := exec.RowsAffected(); num > 0 {
		fmt.Println("==========================================================")
		fmt.Println("add employee  info successful. ")
		fmt.Println("==========================================================")
	}
}
func (e *employee) updated(em *employee) {
	exec, err := dbPool.Exec("UPDATE users set password = ? WHERE id = ?", em.password, em.id)
	if err != nil {
		fmt.Println("==========================================================")
		fmt.Println("updated employee info failed.")
		fmt.Println("==========================================================")
	}
	if num, _ := exec.RowsAffected(); num > 0 {
		fmt.Println("==========================================================")
		fmt.Println("updated employee  info successful. ")
		fmt.Println("==========================================================")
	}
}
func (e *employee) delete(em *employee) {
	exec, err := dbPool.Exec("DELETE  FROM users WHERE id = ?", em.id)
	if err != nil {
		fmt.Println("==========================================================")
		fmt.Println("delete employee info failed.")
		fmt.Println("==========================================================")
	}
	if num, _ := exec.RowsAffected(); num > 0 {
		fmt.Println("==========================================================")
		fmt.Println("delete employee  info successful. ")
		fmt.Println("==========================================================")
	}
}
func (e *employee) find() {
	rows, err := dbPool.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("find employee info failed.")
	}
	defer rows.Close()
	for rows.Next() {
		var em employee
		rows.Scan(&em.id, &em.name, &em.password)
		fmt.Println("==========================================================")
		fmt.Printf("ID:%d 名字:%s 密码:%s \n", em.id, em.name, em.password)
	}
}

func menus() {
	fmt.Println("==========================================================")
	fmt.Println("		   Employee System v1.1      			   ")
	fmt.Println("1.添加员工信息.")
	fmt.Println("2.更新密码信息.")
	fmt.Println("3.删除员工信息.")
	fmt.Println("4.查询所有信息.")
	fmt.Println("==========================================================")
	fmt.Println("请输入对应操作功能的序号:")
	funcNum = func() string {
		readString, _ := reader.ReadString('\n')
		return readString
	}()
	switch funcNum {
	case "1\n":
		es.insert(func() *employee {
			fmt.Println("请输入员工姓名:")
			name, _ := reader.ReadString('\n')
			fmt.Println("请输入员工密码:")
			pwd, _ := reader.ReadString('\n')
			return &employee{strings.Replace(name, "\n", "", -1), pwd, -1}
		}())
	case "2\n":
		// 函数式编程
		es.updated(func() *employee {
			fmt.Println("请输入员工ID:")
			id, _ := reader.ReadString('\n')
			fmt.Println("请输入员工密码:")
			pwd, _ := reader.ReadString('\n')
			return &employee{"", pwd, func() int {
				atoi, _ := strconv.Atoi(strings.Replace(id, "\n", "", -1))
				return atoi
			}()}
		}())
	case "3\n":
		es.delete(func() *employee {
			fmt.Println("请输入员工ID:")
			id, _ := reader.ReadString('\n')
			return &employee{"", "", func() int {
				atoi, _ := strconv.Atoi(strings.Replace(id, "\n", "", -1))
				return atoi
			}()}
		}())
	case "4\n":
		es.find()
	default:
		menus()
	}
	menus()
}

func (dbc *DbConfig) build() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbc.user, dbc.pwd, dbc.addr, dbc.port, dbc.database)
}

func init() {
	dbc := &DbConfig{user: "root", pwd: "root", addr: "192.168.64.2", port: 3306, database: "go_user"}
	dp, err := sql.Open("mysql", dbc.build())
	if err != nil {
		fmt.Println("open database failed.")
		return
	}
	err = dp.Ping()
	if err != nil {
		fmt.Println("database connection failed.")
		return
	}
	dbPool = dp
	dbPool.SetMaxOpenConns(MaxConnNum)
}

func main() {
	menus()
}
