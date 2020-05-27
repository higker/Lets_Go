package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

// GO连接kafka发送消息
// ./kafka-server-start.sh -daemon ../config/server.properties
// ./zookeeper-server-start.sh -daemon ../config/zookeeper.properties
// 基于sarama第三方库开发的kafka client

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test_log"
	msg.Value = sarama.StringEncoder("this is a test log :" + func() string {
		rand.Seed(time.Now().UnixNano())
		str := strconv.Itoa(rand.Int())
		return str
	}())
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"104.215.49.91:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	fmt.Println(msg)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("来了2")
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
