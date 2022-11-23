package config

import (
	"github.com/pelletier/go-toml"
	"os"
)

type Mysql struct {
	Host     string `json:"host" toml:"host"`
	Port     string `json:"port" toml:"port"`
	UserName string `json:"userName" toml:"userName"`
	Password string `json:"password" toml:"password"`
	//TimeOut string `json:"timeOut" toml:"timeOut"`
}

type Config struct {
	Mysql Mysql
}

var Configs = Config{}

func InitConfig(tomlFile string) error {
	file, err := os.ReadFile(tomlFile)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(file, &Configs); err != nil {
		return err
	}
	return nil
}
