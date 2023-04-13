package model

import "github.com/jinzhu/gorm"

// CheckFile 文件校验表
type CheckFile struct {
	FileName   string `json:"fileName"`
	FileSha1   string `json:"fileSha1"`
	ChunkCount int    `json:"chunkCount"`
}

// FileUploadStatus 文件上传状态表
type FileUploadStatus struct {
	gorm.Model
	FileSha1   string `json:"fileSha1"`
	FileName   string `json:"fileName"`
	ChunkCount int    `json:"chunkCount"`
	Status     int    `json:"status"` // 1-上传中，2-上传完成，3-上传失败,4-合并中，5-合并完成，6-合并失败
}

// 文件上传状态枚举表
const (
	Uploading = iota + 1
	UploadFinish
	UploadFailed
	Merging
	MergeFinish
	MergeFailed
)
