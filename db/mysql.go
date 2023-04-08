package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rujiefenfang/filestore-server/config"
	"github.com/rujiefenfang/filestore-server/model"
	"log"
	"os"
)

var dsn string
var mysqlDB *gorm.DB
var err error

func init() {
	mysql := config.Configs.Mysql
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysql.UserName, mysql.Password, mysql.Host, mysql.Port, mysql.DataBase, "10s")

	mysqlDB, err = gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		os.Exit(1)
	}
	//映射表结构
	mapper()

	// 设置最大空闲连接数为 10
	mysqlDB.DB().SetMaxIdleConns(10)

	// 设置最大打开连接数为 100
	mysqlDB.DB().SetMaxOpenConns(100)

}

func GetDB() *gorm.DB {
	return mysqlDB
}
func Close() error {
	return mysqlDB.Close()
}

func mapper() {
	mysqlDB.AutoMigrate(&model.FileMeta{})
}
