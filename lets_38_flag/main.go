// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/15 - 10:04 下午

//flag 包 命令行参数工具
package main

import "flag"

var (
	Ip   string
	Port int
)

func main() {
	flag.StringVar(&Ip, "ip", "127.0.0.1", "请输入IP地址.")
	flag.IntVar(&Port, "port", 8080, "请输入端口地址.")
	flag.Parse() //解析Parse
}
