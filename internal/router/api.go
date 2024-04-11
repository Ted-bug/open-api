package router

import "github.com/gin-gonic/gin"

func initApiRouter(group *gin.RouterGroup) {
	api := group.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
