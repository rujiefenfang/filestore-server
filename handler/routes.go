package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rujiefenfang/filestore-server/auth"
	"github.com/rujiefenfang/filestore-server/handler/files"
	"github.com/rujiefenfang/filestore-server/handler/sign"
	"github.com/rujiefenfang/filestore-server/handler/upload"
	"github.com/rujiefenfang/filestore-server/handler/user"
)

func SetupRoutes(router *gin.Engine) {
	// GET请求的路由路径

	// 用户相关
	group := router.Group("/user")
	group.Use(auth.MyAuth())
	{
		// 用户首页
		group.GET("/home", user.Home)
		// 获取用户信息
		group.POST("/info", user.Info)
		// 查询用户文件列表
		group.POST("/file/query", files.Query)
		// 文件上传
		group.GET("/upload", upload.View)
		group.GET("/get-upload-url", upload.GetUploadUrl)
		group.POST("/upload", upload.Chunks)
		group.POST("/check-file", upload.CheckFile)

	}
	// 用户注册
	router.GET("/user/signup", sign.SignUp)
	// 用户登录
	router.GET("/user/signin", sign.SignIn)
	// 用户登录跳转
	router.POST("/user/signin", sign.PostSignIn)

}
