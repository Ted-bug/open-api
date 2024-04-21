package userservice

import (
	"github.com/Ted-bug/open-api/internal/model"
	"github.com/Ted-bug/open-api/internal/tool/response"
)

func CheckSignWithAk(ak, timestamp, sign string) response.AnyResponse {
	if !model.IsAkSkExist(ak) {
		return response.FailedWithMsg("ak sk error")
	}
	sk, _ := model.GetSk(ak)
	if res, err := model.CheckSign(sk, timestamp, sign); !res || err != nil {
		return response.FailedWithMsg("sign error")
	}
	return nil
}
