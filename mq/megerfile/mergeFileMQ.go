package megerfile

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/rujiefenfang/filestore-server/config"
	"github.com/rujiefenfang/filestore-server/db"
	"github.com/rujiefenfang/filestore-server/model"
	"github.com/rujiefenfang/filestore-server/mq"
	"log"
	"os"
	"strconv"
)

var (
	// RabbitMQURL rabbitmq的链接
	rabbitMQURL string
	// MergeFileProducer 合并文件的生产者
	MergeFileProducer *mq.Producer
	// MergeFileConsumer 合并文件的消费者
	MergeFileConsumer *mq.Consumer
	// err
	err error
)

const (
	rabbitMQExchange     = "filestore_mergefile_exchange"
	rabbitMQExchangeType = "direct"
	rabbitMQQueueName    = "filestore_mergefile_queue"
	rabbitMQKey          = "filestore_mergefile_key"
	rabbitMQDurable      = true
	rabbitMQAutoDelete   = false
)

func init() {
	mqConfig := config.Configs.RabbitMQ
	rabbitMQURL = fmt.Sprintf("amqp://%s:%s@%s:%s/", mqConfig.UserName, mqConfig.Password, mqConfig.Host, mqConfig.Port)
	MergeFileProducer, err = mq.NewProducer(rabbitMQURL, rabbitMQExchange, rabbitMQExchangeType, rabbitMQQueueName, rabbitMQKey, rabbitMQDurable, rabbitMQAutoDelete)
	if err != nil {
		log.Fatalf("failed to connect mergeFile rabbitmq: %v", err)
	}
	MergeFileConsumer, err = mq.NewConsumer(rabbitMQURL, rabbitMQExchange, rabbitMQExchangeType, rabbitMQQueueName, rabbitMQKey)
	if err != nil {
		log.Fatalf("failed to connect mergeFile rabbitmq: %v", err)
	}
}

// MergeFileConsumerHandler 消费者处理函数
func MergeFileConsumerHandler(msg []byte) {
	fileSha1 := string(msg)
	mysqlDB := db.GetDB()
	first := mysqlDB.Where("file_sha1 = ?", fileSha1).First(&model.FileUploadStatus{})
	fileName := first.Value.(*model.FileUploadStatus).FileName
	chunkCount := first.Value.(*model.FileUploadStatus).ChunkCount
	status := first.Value.(*model.FileUploadStatus).Status

	// 判断文件是否正在合并或者已经合并,如果是则不处理
	if status == model.Merging || status == model.MergeFinish {
		return
	}

	// 判断文件是否正在合并或者已经合并
	if status == model.UploadFinish {
		// 合并中
		err = mysqlDB.Transaction(func(tx *gorm.DB) error {
			// 更新文件状态
			return tx.Model(&model.FileUploadStatus{}).Where("file_sha1 = ?", fileSha1).Update("status", model.Merging).Error

		})
		if err != nil {
			log.Fatalf("failed to merge file: %v", err)
		}
		// 合并文件
		err := mergeFile(fileName, chunkCount)
		// 合并失败
		if err != nil {

			errDB := mysqlDB.Transaction(func(tx *gorm.DB) error {
				// 更新文件状态
				return tx.Model(&model.FileUploadStatus{}).Where("file_sha1 = ?", fileSha1).Update("status", model.MergeFailed).Error
			})
			log.Fatalf("failed to merge file: system:%v , mysql:%v", err, errDB)
		}
		// 合并成功
		err = mysqlDB.Transaction(func(tx *gorm.DB) error {
			// 更新文件状态
			return mysqlDB.Model(&model.FileUploadStatus{}).Where("file_sha1 = ?", fileSha1).Update("status", model.MergeFinish).Error
		})
		if err != nil {
			log.Fatalf("failed to merge file: %v", err)
		}
	}

	// 删除切片
	err = deleteChunks(fileName)
	if err != nil {
		log.Fatalf("failed to delete chunks: %v", err)
	}
}

// 合并文件
func mergeFile(fileName string, chunkCount int) error {
	fmt.Println("Merging file:", fileName)

	dstPath := "./uploads/finished/" + fileName
	resPath := "./uploads/" + fileName
	// 创建目标文件
	dstFile, err := os.Create(dstPath + ".tmp")
	if err != nil {
		fmt.Println("Failed to create destination file:", err)
		return err
	}

	// 合并切片
	for i := 0; i < chunkCount; i++ {
		chunkPath := resPath + "/" + strconv.Itoa(i)
		chunkData, err := os.ReadFile(chunkPath)
		if err != nil {
			fmt.Println("Failed to read chunk:", err)
			return err
		}
		dstFile.Write(chunkData)
	}

	// 关闭文件
	defer dstFile.Close()

	// 移动文件
	err = os.Rename(dstPath+".tmp", dstPath)
	if err != nil {
		fmt.Println("Failed to move file:", err)
		return err
	}

	fmt.Println("File", fileName, "merged successfully.")
	return nil
}

// 删除切片
func deleteChunks(fileName string) error {
	resPath := "./uploads/" + fileName
	err := os.RemoveAll(resPath)
	return err
}
