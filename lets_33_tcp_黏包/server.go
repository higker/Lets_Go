// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/11 - 7:14 下午

package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9598")
	if err != nil {
		fmt.Println("create tcp server fail.", err)
		return
	}
	var msg [1024]byte
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("tcp accept fail", err)
			return
		}
		n, _ := accept.Read(msg[:])
		fmt.Println(string(msg[:n]))
	}
}
