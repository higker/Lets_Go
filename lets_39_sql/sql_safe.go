// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/17 - 10:02 下午

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Assets struct {
	id    int
	name  string
	money int
}

var (
	dbsC *sql.DB // 数据库连接池
	urls = "root:root@tcp(192.168.64.2:3306)/assets"
)

func init() {
	dbc, err := sql.Open("mysql", urls)
	if err != nil {
		fmt.Println("database connection fail.")
		return
	}
	err = dbc.Ping()
	if err != nil {
		fmt.Println("username or password error！")
		return
	}
	dbsC = dbc // 返回给连接池
	dbsC.SetMaxOpenConns(10)
	//fmt.Println(db)
}

func main() {
	// sql 注入演示
	notSafe := "'没有这个名字' OR 1=1 #"
	//isSafe := `丁烁`
	sql := fmt.Sprintf("SELECT * FROM user_assets WHERE name = %s", notSafe)
	fmt.Println(sql)
	row := dbsC.QueryRow(sql)
	var a Assets
	err := row.Scan(&a.id, &a.name, &a.money)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
