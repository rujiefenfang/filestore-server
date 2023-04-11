package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/rujiefenfang/filestore-server/config"
)

// GetRedis 连接redis
func GetRedis() *redis.Client {

	redisConfig := config.Configs.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       0,
	})
	return client
}

var redisDB *redis.Client

func init() {
	// 初始化redis连接
	redisDB = GetRedis()
	//最大连接数
	redisDB.Options().MaxRetries = 100
	//最大空闲连接数
	redisDB.Options().MaxRetries = 10
	//连接超时时间
	redisDB.Options().DialTimeout = 100
}
