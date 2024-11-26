package controller

import (
	"net/http"

	"github.com/Ted-bug/open-api/internal/pkg/common"
	"github.com/Ted-bug/open-api/internal/pkg/response"
	shorturlservice "github.com/Ted-bug/open-api/internal/service/short_url_service"
	"github.com/gin-gonic/gin"
)

// 获取短链
func ConvertLurl(c *gin.Context) {
	params := struct {
		Url string `json:"url"`
	}{}
	c.BindJSON(&params)
	if !common.IsValidURL(params.Url) || !common.IsUrlActive(params.Url) {
		c.JSON(http.StatusOK, response.FailedWithMsg("url is not avaliable"))
		return
	}
	if shortUrl, ok := shorturlservice.IsUrlExist(params.Url); ok {
		c.JSON(http.StatusOK, response.SucceedWithData(map[string]any{
			"short_url": shortUrl,
		}))
		return
	}
	if shortUrl, err := shorturlservice.ConvertLurl(params.Url); err != nil {
		c.JSON(http.StatusOK, response.FailedWithMsg(err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, response.SucceedWithData(map[string]any{
			"short_url": shortUrl,
		}))
	}

}

// 解析短链
func RevertSurl(c *gin.Context) {
	short := c.Query("s")
	if short == "" {
		c.JSON(http.StatusOK, response.FailedWithMsg("url is not avaliable"))
		return
	}
	if url, ok := shorturlservice.RevertSurl(short); ok {
		c.Redirect(http.StatusMovedPermanently, url)
		return
	}
	c.JSON(http.StatusOK, response.FailedWithMsg("url is not avaliable"))
}
