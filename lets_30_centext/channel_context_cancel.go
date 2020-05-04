// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/4 - 12:03 下午

package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//会返回一个取消的根节点
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= 3; i++ {
		go worker(ctx, "jobs"+strconv.Itoa(i))
	}
	time.Sleep(5 * time.Second)
	cancelFunc()
}
func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("task stop....")
			return
		default:
			fmt.Println(name, "task running....")
			time.Sleep(time.Millisecond * 200)
		}
	}

}
