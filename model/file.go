package model

type CheckFile struct {
	FileName   string `json:"fileName"`
	FileSha1   string `json:"fileSha1"`
	ChunkCount int    `json:"chunkCount"`
}
