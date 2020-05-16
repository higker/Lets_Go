// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/16 - 4:02 下午

package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	IPAddr, userName, Password string
	Port                       int
}

func ConnectionDB(dbc DBConfig) (*sql.DB, error) {
	//构建数据库连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/go_user", dbc.userName, dbc.Password, dbc.IPAddr, dbc.Port)
	//fmt.Println(dsn)
	// open 这里方法不通过 验证不出来 open方法只是检测数据连接信息是否符合格式！！
	open, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New("open database connection fail.")
	}
	if open.Ping() != nil {
		return nil, errors.New("connection info error！！！")
	}
	return open, nil
}
