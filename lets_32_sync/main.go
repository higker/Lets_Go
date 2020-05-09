// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/9 - 5:20 下午

package main

import (
	"fmt"
	"sync"
)

//go语言中的读写锁和互斥锁演示

var (
	x    = 0
	lock sync.Mutex
	wg   sync.WaitGroup
)

func main() {
	mutex(3, 50000)
	wg.Wait()
	fmt.Println(x)
}

//模拟一个task
func sum(n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		lock.Lock() //加锁
		x += 1
		lock.Unlock() //解锁
	}
}

//模拟一个worker Pool
func mutex(wk int, n int) {
	wg.Add(wk)
	for i := 1; i <= wk; i++ {
		go sum(n)
	}
}
