package simpleLog

import (
	"log"
	"os"
)

var logFile *os.File

func init() {
	err := os.MkdirAll("./log", os.ModePerm)
	if err != nil {
		log.Println("初始化失败：无法创建 log 目录 -", err)
		return
	}

	filePath := "./log/log.txt"
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("初始化失败：无法打开 log.txt 文件 -", err)
		return
	}
	logFile = f
	// 不要关闭logFile，保持打开状态，方便写日志
}

// WriteLog 其他地方写日志
func WriteLog(msg string) {
	if logFile != nil {
		logFile.WriteString(msg + "\n")
	}
}
