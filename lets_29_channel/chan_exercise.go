// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/23 - 2:23 下午

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//channel练习
// 1.启动一个goroutine随机生成100个随机数到ch1中
// 2.然后启动再启动一个goroutine从ch1拿值求平方存放到一个ch2里面
var (
	wg         sync.WaitGroup
	dataChan   chan int
	squareChan chan int
)

func main() {
	//这个50个缓冲区是够用的是因为dataChan在一边生产,squareChan在一边取值
	//dataChan会阻塞生产所有速度肯定没有squareChan快所有50个肯定够用
	//并行！！！！
	dataChan = make(chan int, 50)
	squareChan = make(chan int, 100)
	//wg.Add(2)
	go task1(dataChan)
	go task2(dataChan, squareChan)
	//wg.Wait()
	//变量输出squareChan
	for e := range squareChan {
		fmt.Println(e)
	}
}

//chan<-只写的channel1用来生成随机数
func task1(ch chan<- int) {
	//defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- randNum()
		time.Sleep(randSleep()) //模拟程序休眠几秒钟
	}
	close(ch)
}

//<-chan只读,chan<-只写
func task2(ch <-chan int, cha chan<- int) {
	//defer wg.Done()
	for {
		//这个的意思就是并行携程
		//goroutine1一边随机生成数字 goroutine2一边取ch1数据然后求平方存放到ch2里
		//如果ok等于false说明ch1已经生成的值我们全部取完了
		//那么ch2也要停止
		//可以想象成一个传送带 工厂流水线！！！
		//time.Sleep(randSleep())
		i, ok := <-ch
		if !ok {
			break
		}
		//求平方存储到cha中
		cha <- i * i
	}
	close(cha)
}

//生成随机数
func randNum() int {
	//设置rand随机源
	rand.Seed(time.Now().UnixNano())
	//生成一百以内的随机数
	return rand.Intn(10) + 1
}

//随机休眠几毫秒
func randSleep() time.Duration {
	return time.Millisecond * time.Duration(randNum())
}
