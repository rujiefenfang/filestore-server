package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func myAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取cookie
		cookie := c.GetHeader("cookie")
		//获取cookie
		if strings.Contains(cookie, "filestore_token=") {
			cookie = strings.Split(cookie, "filestore_token=")[1]
			cookie = strings.Split(cookie, ";")[0]
		}
		// 通过jwt解析token

		// 通过token获取用户信息
		// 通过用户信息获取用户权限
	}

}
