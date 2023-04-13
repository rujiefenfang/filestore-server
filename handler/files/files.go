package files

import (
	"github.com/gin-gonic/gin"
	"github.com/rujiefenfang/filestore-server/db"
	"github.com/rujiefenfang/filestore-server/model"
)

// Query 获取文件列表
func Query(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.String(500, "未通过认证，请重新登录")
		return
	}

	mUser := user.(model.User)

	var fileMetas []model.FileMeta

	// 查询用户文件列表
	db.GetDB().Select("file_sha1, file_name, file_size, created_id,updated_id").Where("user_id = ?", mUser.ID).Find(&fileMetas)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": fileMetas,
	})

}
