package middleware

import (
	"net/http"
	"time"

	"github.com/Ted-bug/open-api/internal/tool/response"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 生成一个速率限制中间件
// fillInterval: 填充间隔
// cap: 容量
// quantum: 每个间隔填充的令牌数量
func RateLimitMiddleware(fillInterval time.Duration, cap int64, quantum int64) func(c *gin.Context) {
	if quantum <= 0 {
		quantum = 1
	}
	bucker := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucker.TakeAvailable(1) < 1 {
			c.JSON(http.StatusOK, response.SucceedResponse(nil, 1, "rate limit..."))
			c.Abort()
			return
		}
		c.Next()
	}
}
