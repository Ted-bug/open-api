package logger

import (
	"fmt"
	"time"

	"github.com/Ted-bug/open-api/config"
	"github.com/Ted-bug/open-api/internal/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 记得使用Logger.Sync()刷新缓冲区
var logMap map[string]*zap.Logger

var (
	TYPE_RUN   = "run"
	TYPE_PANIC = "panic"
)

func InitLogger() {
	if logMap == nil {
		logMap = make(map[string]*zap.Logger, 2)
	}
	option := config.AppConfig.Logger
	for _, lname := range option.List {
		switch option.Type {
		case "file":
			// Logger = CreateSyncLogger()
			option.Filename = lname + ".log"
			// 配置了日志切割写入器的Logger
			logMap[lname] = CreateAsyncLogger(option)
		case "command":
			fallthrough
		default:
			// 打印到命令行stdout的Logger
			if tmpLog, err := zap.NewProduction(zap.AddCaller()); err != nil {
				fmt.Println("a logger failed: ", lname, err)
			} else {
				logMap[lname] = tmpLog
			}
		}
	}
	fmt.Println("logger init success")
}

// Close 保证刷写所有日志到磁盘中
// zap日志没有提供关闭句柄的方法
func Close() {
	fmt.Println("close logger...")
	for _, l := range logMap {
		if l != nil {
			l.Sync()
		}
	}
}

// GetLogger 根据名称获取Logger实例
func GetLogger(lname string) *zap.Logger {
	if l, ok := logMap[lname]; ok && l != nil {
		return l
	}
	return nil
}

// CreateSyncLogger 创建一个同步写日志的Logger实例。
//
// 参数:
//
//	option - Logger配置项，用于配置日志的写入方式和其它属性。
//
// 返回值:
//
//	返回一个配置好的*zap.Logger实例，可用于进行日志记录。
func CreateSyncLogger(option config.Logger) *zap.Logger {
	// 1. 根据配置选项获取同步写入器
	w := GetSyncWriter(option)

	// 2. 配置zap核心，使用JSON编码器，设置写入器和日志级别
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // 使用默认的生产环境编码器配置
		w,             // 设置日志写入器为之前获取的同步写入器
		zap.InfoLevel, // 设置日志级别为Info，即默认记录Info及以上的日志
	)

	// 3. 基于配置的核心生成并返回一个具备调用者信息的Logger实例
	logger := zap.New(core, zap.AddCaller())
	return logger
}

// CreateAsyncLogger 创建一个异步写入日志的Logger实例。
//
// 参数:
//
//	option config.Logger - 日志配置项，包含了日志的配置信息，例如异步写入的相关设置。
//
// 返回值:
//
//	*zap.Logger - 返回一个配置好的异步日志记录器。
func CreateAsyncLogger(option config.Logger) *zap.Logger {
	// 获取异步写入日志的Writer
	buffer := GetAsyncWriter(option)
	// 核心日志配置，使用JSON编码器，将日志信息写入到异步Writer中，日志级别为Info
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // 使用JSON格式编码日志
		buffer,
		zap.InfoLevel,
	)
	// 创建并返回一个带调用者信息的Logger实例
	logger := zap.New(core, zap.AddCaller())
	return logger
}

// GetSyncWriter 创建并返回一个同步写入器，用于将日志写入到指定的文件中。
//
// 参数:
//
//	option config.Logger - 包含日志配置信息的结构体，例如日志文件路径、最大尺寸、备份数量和最多保存天数等。
//
// 返回值:
//
//	zapcore.WriteSyncer - 一个实现了 zapcore.WriteSyncer 接口的同步写入器，用于日志的写入操作。
func GetSyncWriter(option config.Logger) zapcore.WriteSyncer {
	// 使用 lumberjack.Logger 实现 zapcore.WriteSyncer 接口，配置日志文件路径、最大尺寸、备份数量和最多保存天数等参数
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   constants.PROJECTPATH + option.Path + option.Filename, // 日志文件完整路径
		MaxSize:    option.MaxSize,                                        // 每个日志文件的最大尺寸（单位：MB）
		MaxBackups: option.MaxBackups,                                     // 保留的日志文件备份数量
		MaxAge:     option.MaxAge,                                         // 日志文件最多保存的天数
	})
	return w
}

// GetAsyncWriter 将同步写入器包装成异步写入器，实现周期性或达到缓冲上限时的数据刷新。
//
// 参数:
//
//	option config.Logger - 日志配置项，用于获取同步写入器。
//
// 返回值:
//
//	*zapcore.BufferedWriteSyncer - 配置好的异步写入器指针。
func GetAsyncWriter(option config.Logger) *zapcore.BufferedWriteSyncer {
	// 根据日志配置获取同步写入器
	w := GetSyncWriter(option)
	// 创建异步写入器，设置缓冲大小为256KB，每5秒刷新一次
	buffer := &zapcore.BufferedWriteSyncer{
		WS:            w,
		Size:          256, // kb
		FlushInterval: time.Second * 5,
	}
	return buffer
}
