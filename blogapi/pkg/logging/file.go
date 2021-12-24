package logging

import (
	"blogapi/pkg/file"
	"blogapi/pkg/setting"
	"fmt"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSettings.RuntimeRootPath, setting.AppSettings.LogSavePath)
	// Output: runtime/logs/
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSettings.LogSaveName,
		time.Now().Format(setting.AppSettings.TimeFormat),
		setting.AppSettings.LogFileExt,
	)
	// Output: log20060102.log
}

func openLogFile(filePath, fileName string) (*os.File, error) {
	wd, _ := os.Getwd()
	src := fmt.Sprintf("%s/%s", wd, filePath)
	if oops := file.CheckPermission(src); oops {
		return nil, fmt.Errorf("无权访问路径: %s", src)
	}
	if err := file.IsNotExistMkdir(src); err != nil {
		return nil, fmt.Errorf("创建文件夹失败: %s", src)
	}

	handle, err := os.OpenFile(fmt.Sprintf("%s%s", src, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败, err: %v", err)
	}
	return handle, nil // 返回文件的句柄
}
