package upload

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/rujiefenfang/filestore-server/db"
	"github.com/rujiefenfang/filestore-server/model"
	"net/http"
	"os"
	"strconv"
	"time"
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

// GetUploadUrl 获取上传链接
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
	// 获取sha1
	fileSha1 := valueMap["fileSha1"][0]

	// 获取文件
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save chunk " + chunkIndex + " of file " + fileName,
		})
		return
	}

	// 保存文件
	err = c.SaveUploadedFile(header, "./uploads/"+fileName+"/"+chunkIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save chunk " + chunkIndex + " of file " + fileName,
		})
		return
	}

	ctx := context.Background()
	redis := db.GetRedis()

	// 将文件块的索引存入redis
	err = redis.RPush(ctx, fileSha1, chunkIndex).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  err.Error(),
		})
		return
	}

	// 设置过期时间为1天
	err = redis.Expire(ctx, fileSha1, time.Hour*24).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// CheckFile 检查文件是否已经上传完成
func CheckFile(c *gin.Context) {

	var finishedChunks []string
	var isExist bool

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

	// 如果不存在
	mysql := db.GetDB()
	// 查看数据库中是否存在记录
	first := mysql.Where("file_sha1 = ?", checkFile.FileSha1).First(&model.FileUploadStatus{})
	rowsAffected := first.RowsAffected
	// 如果存在
	if rowsAffected != 0 {

		isExist = true

		if first.Value.(*model.FileUploadStatus).Status == 2 {
			for i := 0; i < checkFile.ChunkCount; i++ {
				finishedChunks = append(finishedChunks, strconv.Itoa(i))
			}
			c.JSON(http.StatusOK, gin.H{
				"status":         "OK",
				"isExist":        isExist,
				"finishedChunks": finishedChunks,
			})
			return
		}
		ctx := context.Background()
		// 获取redis客户端
		redis := db.GetRedis()
		// 根据文件的sha1值，从数据库中查询记录,如果不存在者设置key=fileSha1，value=[]，并设置过期时间为1天
		mGet := redis.LRange(ctx, checkFile.FileSha1, 0, -1)
		// 如果存在,则获取已经上传的切片
		if len(mGet.Val()) != 0 {
			// 获取已经上传的切片
			finishedChunks = mGet.Val()
		}
		//上传完成
		if len(finishedChunks) == checkFile.ChunkCount {
			//修改上传文件状态，将状态改为2，更新操作要在事务中
			err := mysql.Transaction(func(tx *gorm.DB) error {
				// 更新数据库中的记录
				updateRaw := tx.Model(&model.FileUploadStatus{}).Where("file_sha1 = ?", checkFile.FileSha1).Update("status", 2).RowsAffected
				// 如果更新失败
				if updateRaw == 0 {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status": "Failed",
						"error":  err.Error(),
					})
				}
				return nil
			})
			// 如果更新失败
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Failed",
					"error":  err.Error(),
				})
				return
			}
			// 删除redis中的记录
			redis.Del(ctx, checkFile.FileSha1)

			//TODO:通知合并切片
		}
		// 返回已经上传的切片
		c.JSON(http.StatusOK, gin.H{
			"status":         "OK",
			"isExist":        isExist,
			"finishedChunks": finishedChunks,
		})
		return
	} else {
		// 如果不存在
		// 将文件信息写入数据库
		insertRaw := mysql.Create(&model.FileUploadStatus{
			FileSha1: checkFile.FileSha1,
			FileName: checkFile.FileName,
			Status:   1,
		}).RowsAffected

		// 如果插入失败
		if insertRaw == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Failed",
				"error":  err.Error(),
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":         "OK",
		"isExist":        isExist,
		"finishedChunks": finishedChunks,
	})

}
