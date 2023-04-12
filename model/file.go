package model

import "github.com/jinzhu/gorm"

type CheckFile struct {
	FileName   string `json:"fileName"`
	FileSha1   string `json:"fileSha1"`
	ChunkCount int    `json:"chunkCount"`
}

// FileUploadStatus 文件上传状态表
type FileUploadStatus struct {
	gorm.Model
	FileSha1 string `json:"fileSha1"`
	FileName string `json:"fileName"`
	Status   int    `json:"status"` // 1-上传中，2-上传完成，3-上传失败
}
