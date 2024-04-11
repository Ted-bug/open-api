package router

import "github.com/gin-gonic/gin"

func InitRouter(g *gin.Engine) {
	// Your should add module's router,if you create a new module like app.
	initApiRouter(&g.RouterGroup)
}
