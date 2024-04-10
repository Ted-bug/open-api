package internal

import (
	"fmt"

	"github.com/Ted-bug/open-api/internal/config"
	"github.com/Ted-bug/open-api/internal/tool/logger"
	"github.com/Ted-bug/open-api/internal/tool/mysql"
	"github.com/Ted-bug/open-api/internal/tool/redis"
)

func Run() {

	// 1.启动服务
	if err := config.InitConfig(); err != nil {
		fmt.Printf("Load Config Error: %s\n", err)
		return
	}
	if err := mysql.InitMysql(); err != nil {
		fmt.Printf("Run mysql error: %s\n", err)
		return
	}
	defer mysql.CloseMysql()
	if err := redis.InitRedis(); err != nil {
		fmt.Printf("Run redis error: %s\n", err)
		return
	}
	defer redis.CloseRedis()
	logger.InitLogger()
}
