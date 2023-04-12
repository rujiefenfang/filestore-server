package mq

import (
	"github.com/rujiefenfang/filestore-server/config"
	"log"
)

var (
	// RabbitMQURL rabbitmq的链接
	rabbitMQURL string
	// MergeFileProducer 合并文件的生产者
	MergeFileProducer *Producer
	// MergeFileConsumer 合并文件的消费者
	MergeFileConsumer *Consumer
	// err
	err error
)

const (
	rabbitMQExchange     = "filestore_mergefile_exchange"
	rabbitMQExchangeType = "direct"
	rabbitMQQueueName    = "filestore_mergefile_queue"
	rabbitMQKey          = "filestore_mergefile_key"
	rabbitMQDurable      = true
	rabbitMQAutoDelete   = false
)

func init() {
	mqConfig := config.Configs.RabbitMQ
	rabbitMQURL = mqConfig.Host + ":" + mqConfig.Port
	MergeFileProducer, err = NewProducer(rabbitMQURL, rabbitMQExchange, rabbitMQExchangeType, rabbitMQQueueName, rabbitMQKey, rabbitMQDurable, rabbitMQAutoDelete)
	if err != nil {
		log.Fatalf("failed to connect mergeFile rabbitmq: %v", err)
	}
	MergeFileConsumer, err = NewConsumer(rabbitMQURL, rabbitMQExchange, rabbitMQExchangeType, rabbitMQQueueName, rabbitMQKey)
	if err != nil {
		log.Fatalf("failed to connect mergeFile rabbitmq: %v", err)
	}
	MergeFileConsumer.ConsumeSimple()
}

// 消费者处理函数
func mergeFileConsumerHandler(msg []byte) bool {
	return true
}
