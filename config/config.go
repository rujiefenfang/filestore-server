package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

type mysqlConfig struct {
	Host     string `json:"host" toml:"host"`
	Port     string `json:"port" toml:"port"`
	UserName string `json:"userName" toml:"userName"`
	Password string `json:"password" toml:"password"`
	DataBase string `json:"dataBase" toml:"database"`
	//TimeOut string `json:"timeOut" toml:"timeOut"`
}

type redisConfig struct {
	Host     string `json:"host" toml:"host"`
	Port     string `json:"port" toml:"port"`
	Password string `json:"password" toml:"password"`
}

type rabbitMQConfig struct {
	Host     string `json:"host" toml:"host"`
	Port     string `json:"port" toml:"port"`
	UserName string `json:"userName" toml:"userName"`
	Password string `json:"password" toml:"password"`
}

type config struct {
	Mysql    mysqlConfig
	Redis    redisConfig
	RabbitMQ rabbitMQConfig
}

const tomlFile = "./config/config.toml"

var Configs = config{}

func init() {

	file, err := os.ReadFile(tomlFile)
	if err != nil {
		return
	}

	if err := toml.Unmarshal(file, &Configs); err != nil {
		return
	}
	fmt.Println(Configs)
}
