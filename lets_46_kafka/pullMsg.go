package main

import (
	"testing"

	"github.com/Shopify/sarama"
)

// 消费kafka中的消息

// Kafka consumer object
type KafkaConsumer struct {
	Address []string // kafka server address
	Topic   string   // consumption topic
}

// 不使用值接受者 因为只是内存副本
func (k *KafkaConsumer) Consumption(t *testing.T) {
	consumer, err := sarama.NewConsumer(k.Address, nil) // connection kafka server
	if err != nil {
		t.Log("connection server error:", err)
		return
	}
	partitionsList, err := consumer.Partitions(k.Topic) // set topic
	if err != nil {
		t.Log("set topic error:", err)
		return
	}
	for partition := range partitionsList { // 遍历分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(k.Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			t.Logf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				t.Logf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
		//time.Sleep(time.Second * 5)
	}

}
