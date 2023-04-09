package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rujiefenfang/filestore-server/model"
	"github.com/rujiefenfang/filestore-server/model/vo"
	"os"
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

	//上下文获取用户信息
	user, exists := c.Get("username")
	if !exists {
		c.String(500, "未通过认证，请重新登录")
		return
	}
	mUser := user.(model.User)
	userInfoVO := vo.UserInfoVO{}
	// copy
	userInfoVO.Username = mUser.Username
	userInfoVO.CreatedAt = mUser.CreatedAt

	// 返回用户信息
	c.JSON(200, gin.H{
		"code": 200,
		"user": userInfoVO,
	})

}
