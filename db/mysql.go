package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rujiefenfang/filestore-server/config"
	"log"
)

var dsn string
var DB *gorm.DB
var err error

func InitMysqlPool(mysql config.Mysql) error {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysql.UserName, mysql.Password, mysql.Host, mysql.Port, "springbootbia14", "10s")
	fmt.Println(dsn)
	DB, err = gorm.Open("mysql", dsn)

	if err != nil {
		return err
	}
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatalf("mysql连接池关闭失败;err:%s", err)
			return
		}
	}(DB)
	return nil
}
