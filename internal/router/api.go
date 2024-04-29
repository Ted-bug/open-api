package router

import (
	"github.com/Ted-bug/open-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func initApiRouter(group *gin.RouterGroup) {
	api := group.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello world",
			})
		})
		api.POST("/short-url", controller.ShortUrl)
		api.GET("/jump", controller.ParseUrl)
	}
}
