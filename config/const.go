package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/d1y/yoxi/utils"
)

// WebServerDefaultPort `web` 默认端口
const WebServerDefaultPort int = 3000

var webDistPath = `dist`

var webAssetsPath = `assets`

// Appname `app` 名称
var Appname = `yoxi`

// WebDistPath `web` 打包项目
var WebDistPath = path.Join(utils.Curr(), webDistPath)

// WebAssetsPath `web` 资源文件
var WebAssetsPath = path.Join(utils.Curr(), webAssetsPath)

func createMacPath(x string) string {
	var runFilePath = os.Args[0]
	var run = path.Join(filepath.Dir(runFilePath), "../Resources")
	var r = fmt.Sprintf(`%v/%v`, run, x)
	return r
}

func init() {
	// if !utils.Check(WebDistPath) {

	// }
	if !utils.Check(WebAssetsPath) && runtime.GOOS == "darwin" {
		WebAssetsPath = createMacPath(webAssetsPath)
		WebDistPath = createMacPath(webDistPath)
	}
	// fmt.Println("WebAssetsPath", WebAssetsPath)
	// fmt.Println("WebDistPath", WebDistPath)
}
