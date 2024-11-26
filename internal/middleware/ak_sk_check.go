package middleware

import (
	"net/http"

	"github.com/Ted-bug/open-api/internal/pkg/response"
	userservice "github.com/Ted-bug/open-api/internal/service/user_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AkSkCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := struct {
			Timestamp string `json:"timestamp" validate:"required"`
			Sign      string `json:"sign" validate:"required"`
			Ak        string `json:"ak" validate:"required"`
		}{}
		checker := validator.New()
		params.Timestamp, params.Sign, params.Ak = c.GetHeader("Operation-Time"), c.GetHeader("Sign"), c.GetHeader("AK")
		if err := checker.Struct(params); err != nil {
			c.JSON(http.StatusOK, response.FailedWithMsg(err.Error()))
			c.Abort()
			return
		}
		if sk, err := userservice.GetUserSkByAk(params.Ak); err != nil {
			c.JSON(http.StatusOK, response.FailedWithMsg("ak is not avaliable"))
			return
		} else if ok, err := userservice.CheckSignByAkSk(sk, params.Timestamp, params.Sign); !ok {
			c.JSON(http.StatusOK, response.FailedWithMsg(err.Error()))
			return
		}
		c.Next()
	}
}
