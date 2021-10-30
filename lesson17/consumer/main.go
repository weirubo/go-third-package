package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type myMessageHandler struct {
}

func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	log.Println(string(m.Body))
	return nil
}

func main() {
	// 默认配置信息，如果配置信息被使用之后，将无法修改。
	config := nsq.NewConfig()
	// 创建 Consumer
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}
	consumerStats := consumer.Stats()
	log.Printf("consumerStats:%+v", consumerStats)
	// 给 Consumer 添加处理器，可添加多个，每个 Handler 都运行在单独的 goroutine 中。如果需要在多个 goroutine 中运行同一个 handler，
	// 请使用 consumer.AddConcurrentHandlers(handler Handler, concurrency int)
	// 此外，还可以使用更简洁的方式 consumer.AddHandler(nsq.HandlerFunc(func(message *Message) error{}))。
	consumer.AddHandler(&myMessageHandler{})
	// 连接 nsqd
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
