// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/9 - 9:58 下午

package main

import (
	"fmt"
	"sync"
)

var (
	s      = 0
	lck    sync.Mutex
	wgr    sync.WaitGroup
	result = make(chan int, 10)
	once   sync.Once
)

func main() {
	task()
	for i := range result {
		fmt.Println(i)
	}
}

func task() {
	i := 0
	go add(i)
	go add(i)
}

//通过递归实现for循环  //错误演示不能使用
func add(i int) {
	if i < 100 {
		lck.Lock()
		i++
		s += 1
		result <- s
		lck.Unlock()
		add(i)
	} else {
		once.Do(func() {
			close(result)
		})
	}
}
