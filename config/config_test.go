package config

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	err := InitConfig("./config.toml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Configs)
}
