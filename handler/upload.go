package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rujiefenfang/filestore-server/db"
	"github.com/rujiefenfang/filestore-server/model"
	"github.com/rujiefenfang/filestore-server/util"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetUpload(c *gin.Context) {

	content, err := os.ReadFile("./static/view/index.html")
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	c.Data(200, "text/html; charset=utf-8", content)
}

func PostUpload(c *gin.Context) {

	// 从请求中解析出文件句柄
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "上传失败",
		})
		fmt.Println(err)
		log.Fatal(err)
	}

	// 打开文件句柄
	open, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "上传失败",
		})
		fmt.Println(err)
		log.Fatal(err)
	}
	defer open.Close()

	// 创建文件
	newFile, err := os.Create("./tmp/" + file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "上传失败",
		})
		fmt.Println(err)
		log.Fatal(err)
	}
	defer newFile.Close()

	// 保存文件
	_, err = io.Copy(newFile, open)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "上传失败",
		})
		fmt.Println(err)
		log.Fatal(err)
	}

	// 创建filemeta
	fileMeta := model.FileMeta{
		FileName: file.Filename,
		Location: "./tmp/" + file.Filename,
		FileSize: strconv.FormatInt(file.Size, 10),
		FileSha1: util.FileSha1(newFile),
	}
	db.GetDB().Create(&fileMeta)

	c.JSON(http.StatusOK, gin.H{
		"message":  "上传成功",
		"fileMeta": fileMeta,
	})
}
