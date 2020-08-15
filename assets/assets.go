package assets

import (
	"log"
	"os"
	"path"
)

// Mp3 测试地址
var Mp3 = ""

func init() {
	curr, err := os.Getwd()
	if err != nil {
		log.Fatalln("获取当前路径失败")
	}
	Mp3 = path.Join(curr, "./assets/data/night.mp3")
}
