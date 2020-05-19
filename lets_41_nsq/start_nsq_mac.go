// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/19 - 9:06 下午

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	go func() {
		fmt.Println("star nsqlookupd...")
		exec.Command("bash", "-c", `nsqlookupd`)
	}()
	go func() {
		fmt.Println("star nsqd")
		exec.Command("bash", "-c", "nsqd -lookupd-tcp-address=127.0.0.1:4160")
	}()
	go func() {
		fmt.Println("start nsqadmin")
		exec.Command("bash", "-c", "nsqadmin -lookupd-http-address=127.0.0.1:4161")
	}()
}
