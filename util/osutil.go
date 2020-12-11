package util

import (
	"os"
)

//Exists ファイルの存在チェック
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
