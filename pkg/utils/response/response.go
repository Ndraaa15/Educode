package response

import (
	"errors"

	"github.com/Ndraaa15/Educode/pkg/errsx"
	"github.com/gin-gonic/gin"
)

type response struct {
	Status  status      `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type status struct {
	Code int `json:"code"`
}

func Success(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, response{
		Status: status{
			Code: code,
		},
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, err error, message string, data interface{}) {
	var errsx errsx.Errsx
	if errors.As(err, &errsx) {
		ctx.JSON(code, response{
			Status: status{
				Code: errsx.Code,
			},
			Message: errsx.Message,
			Error:   err.Error(),
			Data:    data,
		})
		return
	} else {
		ctx.JSON(code, response{
			Status: status{
				Code: code,
			},
			Message: message,
			Error:   err.Error(),
			Data:    data,
		})
	}
}
