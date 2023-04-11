package uploads

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

// 合并文件
func mergeFile(fileName string, chunkCount int) {
	fmt.Println("Merging file:", fileName)

	// 创建目标文件
	dstFile, err := os.Create(fileName + "/" + fileName + ".tmp")
	if err != nil {
		fmt.Println("Failed to create destination file:", err)
		return
	}

	// 合并切片
	for i := 0; i < chunkCount; i++ {
		chunkPath := fileName + "/" + strconv.Itoa(i)
		chunkData, err := ioutil.ReadFile(chunkPath)
		if err != nil {
			fmt.Println("Failed to read chunk:", err)
			return
		}
		dstFile.Write(chunkData)
	}

	// 关闭文件
	defer dstFile.Close()

	// 移动文件
	err = os.Rename(fileName+"/"+fileName+".tmp", fileName+"/"+fileName)
	if err != nil {
		fmt.Println("Failed to move file:", err)
		return
	}

	fmt.Println("File", fileName, "merged successfully.")

}

// 计算文件的sha1值
func calculateSHA1(file *os.File) (string, error) {
	h := sha1.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
func TestMergeFile(t *testing.T) {
	//mergeFile("1-1导学【】【更多资源上：  666java.com】.mp4", 14)
	open, err := os.Open("1-1导学【】【更多资源上：  666java.com】.mp4/1-1导学【】【更多资源上：  666java.com】.mp4")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer open.Close()
	fileSha1, err := calculateSHA1(open)
	if err != nil {
		fmt.Println("Failed to calculate sha1:", err)
		return
	}
	fmt.Println(fileSha1 == "8e6a662a9f20f7002ba89207d84b71fbd3427c64")
}
