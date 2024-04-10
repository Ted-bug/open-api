package internal

import (
	"fmt"
	"time"

	"github.com/Ted-bug/open-api/internal/config"
	"github.com/Ted-bug/open-api/internal/middleware"
	"github.com/Ted-bug/open-api/internal/tool/logger"
	"github.com/Ted-bug/open-api/internal/tool/mysql"
	"github.com/Ted-bug/open-api/internal/tool/redis"
	"github.com/gin-gonic/gin"
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

func InitEngine() *gin.Engine {
	gin.SetMode(config.AppConfig.Mode)

	g := gin.New()

	loc, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = loc

	// g.Use(middle.LoggerMiddlerware())
	g.Use(middleware.RecoveryMiddleware())
	// g.Use(middleware.RateLimitMiddleware(10*time.Millisecond, 3000, 2))

	return g
}
