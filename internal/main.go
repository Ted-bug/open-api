package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ted-bug/open-api/config"
	"github.com/Ted-bug/open-api/internal/constants"
	"github.com/Ted-bug/open-api/internal/middleware"
	"github.com/Ted-bug/open-api/internal/pkg/db"
	"github.com/Ted-bug/open-api/internal/pkg/logger"
	"github.com/Ted-bug/open-api/internal/pkg/redis"
	"github.com/Ted-bug/open-api/internal/router"
	"github.com/gin-gonic/gin"
)

func Run() {
	// 1.启动服务
	constants.InitPath()
	if err := config.InitConfig(); err != nil {
		fmt.Printf("load config error: %s\n", err)
		return
	}
	if err := db.InitDB(db.TYPE_MYSQL); err != nil {
		fmt.Printf("run mysql error: %s\n", err)
		return
	}
	defer db.CloseDB()
	if err := redis.InitRedis(); err != nil {
		fmt.Printf("run redis error: %s\n", err)
		return
	}
	defer redis.CloseRedis()
	logger.InitLogger()
	defer logger.Close()

	// 2.获取引擎，配置路由
	g := InitEngine()
	router.InitRouter(g)

	// 3.启动
	// 4.优雅关机
	addr := config.AppConfig.Host + ":" + config.AppConfig.Port
	fmt.Printf("点击访问: http://%s\n", addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: g,
	}
	// 4.1 开启一个goroutine处理请求；否则会一直循环中，无法执行往下的关闭代码
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	// 4.2 创建一个通道监听中断信号
	// kill（syscall.SIGTERM）、kill -2(syscall.SIGINT)监听得到、kill -9监听不到
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // 无信号会阻塞
	fmt.Println("shutdown server ...")
	// 4.3 接收到结束信号，创建5秒超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown: ", err)
	}
	fmt.Println("server exiting")
}

// 初始化引擎
func InitEngine() *gin.Engine {
	gin.SetMode(config.AppConfig.Mode)

	g := gin.New()

	loc, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = loc

	g.Use(middleware.RecoveryMiddlerware())
	g.Use(middleware.LoggerMiddlerware())
	g.Use(middleware.RateLimitMiddleware(10*time.Millisecond, 3000, 2))

	return g
}
