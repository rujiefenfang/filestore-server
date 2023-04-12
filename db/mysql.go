package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rujiefenfang/filestore-server/config"
	"github.com/rujiefenfang/filestore-server/model"
	"log"
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

// TODO sql查询日志
func sqlLog() {
	//mysqlDB.LogMode(true)
	//// log路径
	//file, err := os.Open("./logs/sql.log")
	//if err != nil {
	//	panic(err)
	//}
	//mysqlDB.SetLogger(log.New(file, "\r\n", 0))
}

func mapper() {
	mysqlDB.AutoMigrate(&model.FileMeta{})
	mysqlDB.AutoMigrate(&model.User{})
	mysqlDB.AutoMigrate(&model.FileUploadStatus{})

}
