package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type FileMeta struct {
	gorm.Model
	FileSha1 string `json:"fileSha1,omitempty"`
	FileName string `json:"fileName,omitempty"`
	FileSize string `json:"fileSize,omitempty"`
	Location string `json:"location,omitempty"`
}

var FileMetas map[string]FileMeta

func init() {
	FileMetas = make(map[string]FileMeta)
}

func UpdateFile(file FileMeta) {
	FileMetas[file.FileSha1] = file
}

func GetFile(fileSha1 string) (*FileMeta, error) {
	if file, ok := FileMetas[fileSha1]; ok {
		return &file, nil
	} else {
		return nil, fmt.Errorf("%s", "该文件不存在")
	}

}
