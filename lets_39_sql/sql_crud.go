// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/16 - 9:11 下午

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Go语言连接MySQL 增删改查操作

var (
	db    *sql.DB // 数据库连接池
	dbUrl = "root:root@tcp(192.168.64.2:3306)/go_user"
)

func init() {
	dbc, err := sql.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println("database connection fail.")
		return
	}
	err = dbc.Ping()
	if err != nil {
		fmt.Println("username or password error！")
		return
	}
	db = dbc // 返回给连接池
	db.SetMaxOpenConns(10)
	//fmt.Println(db)
}

type user struct {
	id       int
	name     string
	password string
}

func queryOneById(id int) user {
	var u user
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&u.id, &u.name, &u.password)
	if err != nil {
		fmt.Println(err)
	}
	return u
}
func insert() {

}
func updated() {

}
func delete() {

}

func main() {
	fmt.Println(queryOneById(1))
}
