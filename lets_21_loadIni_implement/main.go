// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Mysql struct {
	Port      int
	IpAddress string
	username  string
	password  string
}

func loadIniFile() *Mysql {
	return nil
}

func main() {
	setValues()
}

//设置值
func setValues() {
	getwd, _ := os.Getwd()
	value := mapValue(openFile(getwd + "/lets_21_loadIni_implement/" + "/mysql.ini"))
	fmt.Println(value)
}

// 通过int文件返回一个map
func mapValue(file *os.File) map[string]interface{} {
	defer file.Close()
	reader := bufio.NewReader(file)
	vm := make(map[string]interface{}, 20)
	for {
		con, err := reader.ReadString('\n')
		split := strings.Split(con, "=")
		vm[split[0]] = split[1]
		if err == io.EOF {
			break
		}
	}
	return vm
}

func openFile(path string) *os.File {
	open, err := os.Open(path)
	if err != nil {
		panic("open ini file fail:" + path)
	}
	return open
}
