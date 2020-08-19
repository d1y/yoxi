package utils

import "os"

// Curr 当前目录
func Curr() string {
	d, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	return d
}

// Check 判断文件/文件夹是否存在
func Check(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
