package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rujiefenfang/filestore-server/db"
	"github.com/rujiefenfang/filestore-server/model"
	"strings"
)

// Secret 密钥
var Secret = []byte("filestore-server")

// MyAuth 验证用户是否登录
func MyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取cookie
		cookie := c.GetHeader("cookie")
		if strings.Contains(cookie, "filestore_token=") {
			cookie = strings.Split(cookie, "filestore_token=")[1]
			cookie = strings.Split(cookie, ";")[0]
		}
		// 通过jwt解析token
		claims, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return Secret, nil
		})
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "token过期，请重新登录",
			})
		}

		// 通过token获取用户信息
		mapClaims := claims.Claims.(*jwt.MapClaims)
		username := (*mapClaims)["username"].(string)
		dbUser := model.User{}
		first := db.GetDB().Where("username = ?", username).First(&dbUser)
		//用户不存在
		if first.Error != nil {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "用户不存在，请重新登录",
			})
		}
		// 通过用户信息获取用户权限

		// 将用户信息保存到上下文中
		c.Set("username", dbUser)
		return

	}

}
