package util

import (
	"fmt"
	"os"
	"strconv"
)

// MergeFile 合并文件
func MergeFile(fileName string, chunkCount int) {
	fmt.Println("Merging file:", fileName)
	root := "./uploads/"

	path := root + fileName + "/" + fileName
	// 创建目标文件
	dstFile, err := os.Create(path + ".tmp")
	if err != nil {
		fmt.Println("Failed to create destination file:", err)
		return
	}

	// 合并切片
	for i := 0; i < chunkCount; i++ {
		chunkPath := fileName + "/" + strconv.Itoa(i)
		chunkData, err := os.ReadFile(chunkPath)
		if err != nil {
			fmt.Println("Failed to read chunk:", err)
			return
		}
		dstFile.Write(chunkData)
	}

	// 关闭文件
	defer dstFile.Close()

	// 移动文件
	err = os.Rename(root+".tmp", root)
	if err != nil {
		fmt.Println("Failed to move file:", err)
		return
	}

	fmt.Println("File", fileName, "merged successfully")
}
