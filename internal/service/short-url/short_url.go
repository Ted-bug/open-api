package shorturl

import (
	"github.com/Ted-bug/open-api/internal/model"
	"github.com/Ted-bug/open-api/internal/tool/response"
)

func CreateShortUrl(Url string) response.AnyResponse {
	var short string
	var err error
	var ok bool
	// 1.检查redis
	// 2.检查mysql
	// 3.生成短链接
	if short, ok = model.IsUrlExist(Url); ok {
		return response.SucceedWithData(map[string]any{"short": short})
	}
	if short, err = model.CreateShortUrl(Url); err != nil {
		return response.FailedWithMsg(err.Error())
	}
	return response.SucceedWithData(map[string]any{
		"short": short,
	})
}
