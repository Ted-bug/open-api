package logger

import (
	"github.com/Ted-bug/open-api/config"
	"github.com/Ted-bug/open-api/internal/pkg/logger/zaplog"
	"go.uber.org/zap"
)

const (
	PkgZap = "zap"
)

func InitLogger() {
	option := config.AppConfig.Logger
	switch option.Package {
	case PkgZap:
		zaplog.InitZap()
	}
}

func Close() {
	switch config.AppConfig.Logger.Package {
	case PkgZap:
		zaplog.CloseZap()
	}
}

func GetZapLogger(lname string) *zap.Logger {
	return zaplog.GetLogger(lname)
}
