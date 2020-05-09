// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package main

import (
	"fmt"
	"time"
)

func work(wid int, jobs <-chan int, result chan<- int, state chan<- struct{}) {
	for job := range jobs {
		result <- job * job
		state <- struct{}{} //状态通道
		fmt.Println("wid:", wid, "over:", job)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	state := make(chan struct{}, 100) //状态通道
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go work(w, jobs, results, state)
	}
	// 5个任务
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs) //关闭任务通道
	}()
	go func() {
		for i := 0; i < 5; i++ {
			<-state //取不到就会阻塞 如果全部取完了就说明task完成了关闭计算结果通道
		}
		close(results)
	}()
	for re := range results {
		fmt.Println(re)
	}
}
