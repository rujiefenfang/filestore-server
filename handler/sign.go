package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rujiefenfang/filestore-server/model"
	"os"
)

var Secret = []byte("filestore-server")

func SignIn(c *gin.Context) {
	content, err := os.ReadFile("./static/view/signin.html")
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	c.Data(200, "text/html; charset=utf-8", content)
}
func PostSignIn(c *gin.Context) {
	rawData, err := c.GetRawData()
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	user := model.User{}
	err = json.Unmarshal(rawData, &user)
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	fmt.Println(user)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": base64.StdEncoding.EncodeToString([]byte(user.Password)),
	}).SignedString(Secret)
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
		"token":    token,
		"location": "http://localhost:8080/user",
	})
}
func SignUp(c *gin.Context) {
	content, err := os.ReadFile("./static/view/signup.html")
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	c.Data(200, "text/html; charset=utf-8", content)
}
