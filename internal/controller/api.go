package controller

import (
	"net/http"

	shorturlservice "github.com/Ted-bug/open-api/internal/service/short_url_service"
	userservice "github.com/Ted-bug/open-api/internal/service/user_service"
	"github.com/Ted-bug/open-api/internal/tool/response"
	"github.com/gin-gonic/gin"
)

// 获取短链
func ShortUrl(c *gin.Context) {
	param := struct {
		Url       string `json:"url"`
		Timestamp string `json:"timestamp"`
		Sign      string `json:"sign"`
		Ak        string `json:"ak"`
	}{}
	c.BindJSON(&param)
	param.Timestamp, param.Sign, param.Ak = c.GetHeader("Operation-Time"), c.GetHeader("Sign"), c.GetHeader("AK")
	if param.Timestamp == "" || param.Sign == "" || param.Ak == "" || param.Url == "" {
		c.JSON(http.StatusOK, response.FailedWithMsg("header or param is not avaliable"))
		return
	}
	var (
		shortUrl string
		sk       string
		err      error
		ok       bool
	)
	if sk, err = userservice.GetUserSkByAk(param.Ak); err != nil {
		c.JSON(http.StatusOK, response.FailedWithMsg("ak is not avaliable"))
		return
	}
	if ok, err = userservice.CheckSignByAkSk(sk, param.Timestamp, param.Sign); !ok {
		c.JSON(http.StatusOK, response.FailedWithMsg(err.Error()))
		return
	}
	if shortUrl, ok = shorturlservice.IsUrlExist(param.Url); ok {
		c.JSON(http.StatusOK, response.SucceedWithData(map[string]any{
			"short_url": shortUrl,
		}))
		return
	}
	if shortUrl, err = shorturlservice.CreateShortUrl(param.Url); err != nil {
		c.JSON(http.StatusOK, response.FailedWithMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SucceedWithData(map[string]any{
		"short_url": shortUrl,
	}))
}
