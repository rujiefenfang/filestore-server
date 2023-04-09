package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rujiefenfang/filestore-server/db"
	"github.com/rujiefenfang/filestore-server/handler"
	"github.com/rujiefenfang/filestore-server/logs"
	"log"
	"os"
)

func main() {

	// 初始化log文件
	{
		if err := logs.LogInit(); err != nil {
			fmt.Printf("log文件初始化失败; err:%s\n", err)
			os.Exit(1)
		}

	}
	// 初始化mysql
	mysqlDB := db.GetDB()

	// 关闭mysql
	defer func(mysqlDB *gorm.DB) {
		err := mysqlDB.Close()
		if err != nil {
			log.Fatalf("mysql关闭失败，%s", err)
		}
	}(mysqlDB)

	// 初始化路由
	r := gin.Default()
	// 静态资源
	r.Static("templates", "./static")
	// 加载路由
	handler.SetupRoutes(r)
	// 启动服务
	r.Run(":8081")

}
