package middleware

import (
	"github.com/Ted-bug/open-api/internal/tool/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddlerware() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(logger.Logger, true)
}
