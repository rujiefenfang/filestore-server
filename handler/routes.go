package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rujiefenfang/filestore-server/auth"
)

func SetupRoutes(router *gin.Engine) {
	// GET请求的路由路径
	// 文件
	// 上传文件
	router.GET("/file/upload", GetUpload)
	router.POST("/file/upload", PostUpload)

	// 用户相关
	group := router.Group("/user")
	group.Use(auth.MyAuth())
	{
		// 用户首页
		group.GET("/home", Home)
		// 获取用户信息
		group.POST("/info", UserInfo)
	}
	// 用户注册
	router.GET("/user/signup", SignUp)
	// 用户登录
	router.GET("/user/signin", SignIn)
	// 用户登录跳转
	router.POST("/user/signin", PostSignIn)

}
