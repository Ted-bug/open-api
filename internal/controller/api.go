package controller

import (
	shorturl "github.com/Ted-bug/open-api/internal/service/short-url"
	userservice "github.com/Ted-bug/open-api/internal/service/user-service"
	"github.com/Ted-bug/open-api/internal/tool/response"
	"github.com/gin-gonic/gin"
)

// 获取短链
func ShortUrl(c *gin.Context) {
	param := struct {
		Url string `json:"url"`
	}{}
	timestamp := c.GetHeader("Operation-Time")
	sign := c.GetHeader("Sign")
	ak := c.GetHeader("AK")
	c.BindJSON(&param)
	if timestamp == "" || sign == "" || ak == "" || param.Url == "" {
		c.JSON(200, response.FailedWithMsg("header or param is not avaliable"))
		return
	}
	if res := userservice.CheckSignWithAk(ak, timestamp, sign); res != nil {
		c.JSON(200, userservice.CheckSignWithAk(ak, timestamp, sign))
	}
	c.JSON(200, shorturl.CreateShortUrl(param.Url))
}
