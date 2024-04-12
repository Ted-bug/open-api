package constants

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// 定义使用到的常量
var (
	PROJECTPATH = ""
	CONFPATH    = "/config"        // 配置文件路径
	STATIC      = "/public/static" // 静态文件路径
	TEMPLATES   = "/templates/"    // 模板文件路径
	LOGPATH     = "/logs/"         // 日志路径
)

var (
	LOGIN_TOKEN_SALT = []byte("www.ted.com")
)

func InitPath() {
	PROJECTPATH = GetProjectPath()
	CONFPATH = PROJECTPATH + CONFPATH
	STATIC = PROJECTPATH + STATIC
	TEMPLATES = PROJECTPATH + TEMPLATES
	LOGPATH = PROJECTPATH + LOGPATH
}

// 获取项目根路径
func GetProjectPath() string {
	// 1.获取exe的生成路径
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	execPath, _ = filepath.EvalSymlinks(filepath.Dir(execPath))

	// 2.获取期待运行的路径（即项目在哪，期待的路径就在哪，但好像会被build进exe中）
	var runPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		runPath = path.Dir(path.Dir(path.Dir(filename))) // !!需要根据此函数的go文件所处的项目相对位置调节Dir层数
	}

	// 3.获取Temp路径（好像linux没有，win有，故而判断该用哪个路径）
	tempPath := os.Getenv("TEMP")
	if tempPath == "" {
		tempPath = os.Getenv("TMP")
	}
	tempPath, _ = filepath.EvalSymlinks(tempPath)

	// 4.识别最终路径
	if strings.Contains(execPath, tempPath) {
		return runPath
	}
	return execPath
}
