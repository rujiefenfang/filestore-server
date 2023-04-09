package logs

import (
	"log"
	"os"
	"time"
)

type LogFile struct {
	fp string
}

func (lf *LogFile) Write(p []byte) (int, error) {
	file, err := os.OpenFile(lf.fp, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		return -1, err
	}
	return file.Write(p)
}

func LogInit() error {
	lf := LogFile{
		fp: "./logs/server.log",
	}
	log.SetOutput(&lf)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	nowTime := time.Now().Format("2006/01/02 15:04:05")

	top := nowTime + " ServerLogs" + " =============================\n"

	_, err := lf.Write([]byte(top))

	return err
}
