package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

// 生产者
func main() {
	// 默认配置信息
	config := nsq.NewConfig()

	// 创建生产者
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	// 验证生成者连接是否成功
	err = producer.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// 返回生产者地址
	producerAddr := producer.String()
	log.Printf("producerAddr:%v", producerAddr)
	messageBody := []byte("hello")
	topicName := "topic"
	// 同步发送消息到指定 topic
	err = producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}
	producer.Stop()
}
