package middleware

import (
	"github.com/Ted-bug/open-api/internal/pkg/logger"
	"github.com/Ted-bug/open-api/internal/pkg/logger/zaplog"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddlerware() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(logger.GetZapLogger(zaplog.LogPanic), true)
}
