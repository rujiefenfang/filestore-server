package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rujiefenfang/filestore-server/handler"
	"github.com/rujiefenfang/filestore-server/logs"
	"github.com/rujiefenfang/filestore-server/mq/megerfile"
	"log"
)

func main() {

	{
		// 初始化log文件
		if err := logs.LogInit(); err != nil {
			log.Fatalf("failed to init log: %v", err)
		}
		// 开启rabbitmq消费者
		go func() {
			err := megerfile.MergeFileConsumer.ConsumeMessageHandler(megerfile.MergeFileConsumerHandler)
			if err != nil {
				log.Fatalf("failed to consume mergeFile rabbitmq: %v", err)
			}
		}()

	}

	// 初始化路由
	r := gin.Default()
	// 静态资源
	r.Static("templates", "./static")
	// 加载路由
	handler.SetupRoutes(r)
	// 启动服务
	r.Run(":8081")

}
