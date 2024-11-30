package middleware

import (
	"time"

	"github.com/Ted-bug/open-api/internal/pkg/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerMiddlerware() gin.HandlerFunc {
	return ginzap.GinzapWithConfig(logger.GetLogger(logger.TYPE_RUN), &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		Context: func(c *gin.Context) []zapcore.Field {
			fields := []zapcore.Field{}
			traceId := c.Request.Header.Get("X-Trace-Id")
			if traceId == "" {
				traceId = uuid.NewV4().String()
				c.Request.Header.Set("X-Trace-Id", traceId)
			}
			fields = append(fields, zap.String("traceId", traceId))
			return fields
		},
	})
}
