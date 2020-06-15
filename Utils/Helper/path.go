package Helper

import (
	"os"
)

//path是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

//创建path
func MkdirPath(path string) {
	if exist := PathExist(path); exist == false {
		_ = os.MkdirAll(path, os.ModePerm)
	}
}
