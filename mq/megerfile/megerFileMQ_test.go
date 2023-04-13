package megerfile

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"github.com/rujiefenfang/filestore-server/mq"
	"log"
	"os"
	"testing"
	"time"
)

type rabbitMqConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	UserName string `toml:"userName"`
	Password string `toml:"password"`
}

func readConfig() rabbitMqConfig {
	var rabbitMqConfig rabbitMqConfig
	file, err := os.ReadFile("./configMQTest.toml")
	if err != nil {
		return rabbitMqConfig
	}
	if err := toml.Unmarshal(file, &rabbitMqConfig); err != nil {
		return rabbitMqConfig
	}
	return rabbitMqConfig
}

func TestReadConfig(t *testing.T) {
	readConfig()
}

func TestMQ(t *testing.T) {
	mqConfig := readConfig()
	rabbitMQURL = fmt.Sprintf("amqp://%s:%s@%s:%s/", mqConfig.UserName, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	mergeFileProducer, err := mq.NewProducer(rabbitMQURL, rabbitMQExchange, rabbitMQExchangeType, rabbitMQQueueName, rabbitMQKey, rabbitMQDurable, rabbitMQAutoDelete)
	if err != nil {
		log.Fatalf("failed to connect mergeFile rabbitmq: %v", err)
	}
	mergeFileConsumer, err := mq.NewConsumer(rabbitMQURL, rabbitMQExchange, rabbitMQExchangeType, rabbitMQQueueName, rabbitMQKey)
	if err != nil {
		log.Fatalf("failed to connect mergeFile rabbitmq: %v", err)
	}

	for i := 0; i < 10; i++ {
		mergeFileProducer.PublishSimple(fmt.Sprintf("hello world %d", i))
		time.Sleep(time.Second * 2)
	}
	go func() {
		err := mergeFileConsumer.ConsumeMessageHandler(MergeFileConsumerHandler)
		if err != nil {
			log.Fatalf("failed to consume message: %v", err)
		}
	}()
	for i := 0; i < 10; i++ {
		mergeFileProducer.PublishSimple(fmt.Sprintf("hello world %d", i))
		time.Sleep(time.Second * 2)
	}
	select {}

}
