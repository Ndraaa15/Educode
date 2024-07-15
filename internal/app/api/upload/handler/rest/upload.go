package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/Ndraaa15/Educode/middleware"
	"github.com/Ndraaa15/Educode/pkg/cloudinary"
	"github.com/Ndraaa15/Educode/pkg/logx"
	"github.com/Ndraaa15/Educode/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	cld  cloudinary.CloudinaryItf
	logx *logx.Logx
}

func NewUploadHandler(cld cloudinary.CloudinaryItf, logx *logx.Logx) *UploadHandler {
	return &UploadHandler{
		cld:  cld,
		logx: logx,
	}
}

func (h *UploadHandler) Start(srv *gin.RouterGroup) {
	srv.POST("/upload", middleware.ValidateJWTToken("student", "teacher"), h.UploadFile)

}

func (h *UploadHandler) UploadFile(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	var (
		errx error
		code int = http.StatusCreated
		data interface{}
	)

	defer func() {
		if errx != nil {
			h.logx.ErrorLogger(errx)
			response.Error(ctx, code, errx, http.StatusText(code), nil)
			return
		}
		response.Success(ctx, code, http.StatusText(http.StatusCreated), data)
	}()

	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		errx = err
		code = http.StatusBadRequest
		return
	}

	url, err := h.cld.UploadFile(c, file)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = url
}
