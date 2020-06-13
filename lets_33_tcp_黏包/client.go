// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/11 - 7:58 下午

package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:9598")
	if err != nil {
		fmt.Println("connection tcp server fail.", err)
	}
	var buffs [1024]byte

	for  {
		dial.Write([]byte("Hello !!" + time.Now().Format("2006-01-02 15:04:05.0000") + "\n"))
		read, err := dial.Read(buffs[:])
		fmt.Println(string(buffs[:read]))
		if err == io.EOF {
			fmt.Println("服务器已经关闭你的连接！！")
			break
		}
		continue //停止当前的代码 去执行下一次的 2次  1—>fmt -> nil -> nil -> EOF -> for -> continue
		time.Sleep(time.Second * 2)
	}

}
