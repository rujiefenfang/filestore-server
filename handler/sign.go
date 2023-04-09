package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rujiefenfang/filestore-server/db"
	"github.com/rujiefenfang/filestore-server/model"
	"golang.org/x/crypto/bcrypt"
	"os"
)

// Secret 密钥
var Secret = []byte("filestore-server")

// SignIn 登录页面
func SignIn(c *gin.Context) {
	//页面
	content, err := os.ReadFile("./static/view/signin.html")
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	c.Data(200, "text/html; charset=utf-8", content)
}

// PostSignIn 登录接口
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
	dbUser := model.User{}
	// 查询用户
	db.GetDB().Where("username = ?", user.Username).First(&dbUser)

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 通过加密盐生成token，并过滤掉密码
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	}).SignedString(Secret)
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}

	// 返回token
	c.JSON(200, gin.H{
		"username": user.Username,
		"token":    token,
		"location": "http://localhost:8081/user/home",
	})
}

// SignUp 注册页面
func SignUp(c *gin.Context) {
	content, err := os.ReadFile("./static/view/signup.html")
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	c.Data(200, "text/html; charset=utf-8", content)
}
