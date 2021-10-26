package logging

import (
	"blogapi/pkg/setting"
	"fmt"
	"log"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSettings.LogSavePath)
}

func getLogFileFullPath() string {
	path := getLogFilePath()
	fileName := fmt.Sprintf("%s%s.%s", setting.AppSettings.LogSaveName,
		time.Now().Format(setting.AppSettings.TimeFormat), setting.AppSettings.LogFileExt)
	return fmt.Sprintf("%s%s", path, fileName)
	// Output: runtime/logs/log20060102.log
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkLogDir() // 创建日志文件所属的目录
	case os.IsPermission(err):
		log.Fatalf("permission: %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to open file: %v, err: %v", filePath, err)
	}
	return handle // 返回文件的句柄
}

func mkLogDir() {
	dir, _ := os.Getwd()
	if err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm); err != nil {
		panic(err) // TODO
	}
}
