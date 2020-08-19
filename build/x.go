package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"

	"github.com/d1y/macapp"
	"github.com/d1y/yoxi/config"
)

func main() {
	buildMacApp(config.Appname)
}

// 创建 `macapp`
func buildMacApp(appname string) {
	var distPath = path.Join(curr(), "../dist")

	fmt.Println("当前打包目录: ", distPath)

	var Yoxi = macapp.Create(macapp.AppConfig{
		AppName: appname,
		AppPath: distPath,
	})

	// create icon file
	var iconFile = path.Join(curr(), "../logo/logo.png")
	fmt.Println("图标路径", iconFile)
	_, err := Yoxi.SetIcon(iconFile)
	if err != nil {
		log.Fatalln("创建图标失败") // create app icon is errr
	}

	var appPath = path.Join(distPath, fmt.Sprintf("%v.app", appname))

	// create bin file
	var offsetFilePath = fmt.Sprintf("%v.app/Contents/MacOS/%v", appname, appname)
	var outFilePath = path.Join(distPath, offsetFilePath)
	r := exec.Command("go", "build", "-o", outFilePath, "..")
	_, err = r.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("打包二进制文件成功")

	var distWebPath = path.Join(curr(), "../../yoxi_web/dist")
	if !check(distWebPath) {
		log.Fatalln("web项目未打包..")
	}
	var outputWebPath = path.Join(appPath, "./Contents/Resources/dist")
	fmt.Println("web打包目录: ", outputWebPath)
	e := CopyDirectory(distWebPath, outputWebPath)
	if err != nil {
		log.Fatalln(e)
	}

	var outputWebAssets = path.Join(appPath, "./Contents/Resources/assets")
	ensureDir(outputWebAssets)
	fmt.Println("打包静态资源: ", outputWebAssets)

	var assetsResultPath = path.Join(curr(), "../../yoxi_data/results")
	fmt.Println("静态资源目录: ", assetsResultPath)
	e = CopyDirectory(assetsResultPath, outputWebAssets)
	if err != nil {
		log.Fatalln(e)
	}

}

func curr() string {
	d, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	return d
}

// 判断文件/文件夹是否存在
func check(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// copy 复制文件
func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// CopyDirectory https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func CopyDirectory(scrDir, dest string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// Copy https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

// // Exists https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateIfNotExists https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

// CopySymLink https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}

// ensureDir 自动创建文件
func ensureDir(fileName string) {
	dirName := fileName
	// dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, 0755)
		if merr != nil {
			panic(merr)
		}
	}
}
