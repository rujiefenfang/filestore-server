package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
)

func Home(c *gin.Context) {

	content, err := os.ReadFile("./static/view/home.html")
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	c.Data(200, "text/html; charset=utf-8", content)
}

func UserInfo(c *gin.Context) {
	cookie := c.GetHeader("cookie")
	//获取cookie
	if strings.Contains(cookie, "filestore_token=") {
		cookie = strings.Split(cookie, "filestore_token=")[1]
		cookie = strings.Split(cookie, ";")[0]
	}
	// 通过jwt解析token
	claims, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	// 通过token获取用户信息
	claimsMap := *claims.Claims.(*jwt.MapClaims)
	fmt.Println(claimsMap)
	username := claimsMap["username"].(string)
	password := claimsMap["password"].(string)
	bytes, err := base64.StdEncoding.DecodeString(password)

	decodeString := string(bytes)

	fmt.Println(decodeString)
	fmt.Println(username, password)
	c.JSON(200, gin.H{
		"code":   200,
		"cookie": cookie,
	})

}
