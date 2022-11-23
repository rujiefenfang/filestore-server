package main

import (
	"github.com/rujiefenfang/filestore-server/config"
	"github.com/rujiefenfang/filestore-server/db"
	"log"
)

func main() {

	//init config
	if err := config.InitConfig("./config/config.toml"); err != nil {
		log.Fatalf("config.InitConfig: 配置文件加载失败;err:%s", err)
		return
	}

	//init mysqlDB
	if err := db.InitMysqlPool(config.Configs.Mysql); err != nil {
		log.Fatalf("db.InitMysqlPool: mysql连接池创建失败;err:%s", err)
		return
	}

}
