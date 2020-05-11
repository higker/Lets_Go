// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/11 - 8:52 下午

package main

import (
	"fmt"
	"net"
)

const (
	PORT = 9598
)

func main() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: PORT})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer udp.Close()

	for {
		var data [1024]byte
		n, addr, err := udp.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = udp.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
