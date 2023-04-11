package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pelletier/go-toml/v2"
	"os"
	"testing"
)

func readConfig() {
	file, err := os.ReadFile("./config.toml")
	if err != nil {
		return
	}

	if err := toml.Unmarshal(file, &Configs); err != nil {
		return
	}

	fmt.Println(Configs)
}

func TestName(t *testing.T) {
	readConfig()
}

func TestRedis(t *testing.T) {
	readConfig()
	redisConfig := Configs.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       0,
	})
	fmt.Println(client)
}
