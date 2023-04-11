package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rujiefenfang/filestore-server/model"
	"net/http"
	"os"
	"strconv"
)

// View 上传界面
func View(c *gin.Context) {
	content, err := os.ReadFile("./static/view/temp.html")
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	c.Data(200, "text/html; charset=utf-8", content)
}

// GetUploadUrl 获取上传链接 TODO: 添加redis缓存已经上传的文件的块数
func GetUploadUrl(c *gin.Context) {
	fileName := c.Query("fileName")
	chunkCount, _ := strconv.Atoi(c.Query("chunkCount"))

	fmt.Println("Getting upload URL for file:", fileName)

	// 创建目录
	err := os.MkdirAll("./uploads/"+fileName, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create directory for file: " + fileName,
		})
		return
	}

	// 返回上传链接
	uploadUrl := "/user/upload?fileName=" + fileName + "&chunkCount=" + strconv.Itoa(chunkCount)
	c.JSON(http.StatusOK, gin.H{
		"uploadUrl": uploadUrl,
	})
}

// Chunks 上传切片
func Chunks(c *gin.Context) {
	fileName := c.Query("fileName")

	form, _ := c.MultipartForm()
	// 获取表单数据
	valueMap := form.Value
	// 获取文件块数
	chunkIndex := valueMap["chunkIndex"][0]

	fmt.Println("Uploading chunk", chunkIndex, "of file", fileName)

	// 保存切片
	_, header, _ := c.Request.FormFile("file")
	err := c.SaveUploadedFile(header, "./uploads/"+fileName+"/"+chunkIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save chunk " + chunkIndex + " of file " + fileName,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// CheckFile 检查文件是否已经上传
func CheckFile(c *gin.Context) {

	checkFile := model.CheckFile{}
	// 检查请求参数
	err := c.ShouldBindBodyWith(&checkFile, binding.JSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  err.Error(),
		})
		return
	}

	//TODO: 检查文件是否已经上传
	//根据文件的sha1值，从数据库中查询记录
	var finishedChunks []string
	var isExist bool

	c.JSON(http.StatusOK, gin.H{
		"status":         "OK",
		"isExist":        isExist,
		"finishedChunks": finishedChunks,
	})
}
