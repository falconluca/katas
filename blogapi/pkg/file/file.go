package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	c, err := ioutil.ReadAll(f)
	return len(c), err
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist 文件是否存在
func CheckNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// CheckPermission 文件权限 FIXME Deprecated
func CheckPermission(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)
}

func HasPermission(path string) bool {
	return !CheckPermission(path)
}

// IsNotExistMkdir 不存在则创建文件夹
func IsNotExistMkdir(path string) error {
	if exist := CheckNotExist(path); !exist {
		return nil
	}

	// 不存在
	return Mkdir(path)
}

// Mkdir 创建文件夹
func Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}
