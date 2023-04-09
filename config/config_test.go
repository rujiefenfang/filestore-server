package config

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestName(t *testing.T) {
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	fmt.Println(string(password))
}
