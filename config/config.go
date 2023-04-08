package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

type Mysql struct {
	Host     string `json:"host" toml:"host"`
	Port     string `json:"port" toml:"port"`
	UserName string `json:"userName" toml:"userName"`
	Password string `json:"password" toml:"password"`
	DataBase string `json:"dataBase" toml:"database"`
	//TimeOut string `json:"timeOut" toml:"timeOut"`
}

type Config struct {
	Mysql Mysql
}

const tomlFile = "./config/config.toml"

var Configs = Config{}

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
