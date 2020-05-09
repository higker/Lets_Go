// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	xns = 0
	lk  sync.Mutex
	Wgs sync.WaitGroup
)

//go语言中的读写互斥锁  应用场景就是读大于写的时候！！！！！
func main() {
	now := time.Now()
	//写大于读 因为有时候启动goroutine也要时间 如果读放前面他就可能读的是空的
	for i := 0; i < 10; i++ {
		Wgs.Add(1)
		go write1()
	}
	for i := 0; i < 10000; i++ {
		Wgs.Add(1)
		go read1()
	}
	////写大于读 因为有时候启动goroutine也要时间 如果读放前面他就可能读的是空的
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go write()
	//}
	Wgs.Wait()
	fmt.Println(time.Now().Sub(now))
}

//func read() {
//	defer Wg.Done()
//	//我在读的时候就不要做其他操作了防止错误
//	rwl.RLock()
//	fmt.Println(xn)
//	rwl.RUnlock()
//	time.Sleep(2 * time.Millisecond)
//}
//func write() {
//	defer Wg.Done()
//	//我在写的时候就不要做其他操作了防止错误
//	rwl.Lock()
//	xn = xn + 1
//	rwl.Unlock()
//	time.Sleep(5 * time.Millisecond)
//}
func read1() {
	defer Wgs.Done()
	//我在读的时候就不要做其他操作了防止错误
	lk.Lock()
	fmt.Println(xns)
	lk.Unlock()
	time.Sleep(2 * time.Millisecond)
}
func write1() {
	defer Wgs.Done()
	//我在写的时候就不要做其他操作了防止错误
	lk.Lock()
	xns = xns + 1
	lk.Unlock()
	time.Sleep(5 * time.Millisecond)
}
