package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Ndraaa15/IQuest/internal/app/api/video/service"
	"github.com/Ndraaa15/IQuest/pkg/logx"
	"github.com/Ndraaa15/IQuest/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	video service.IVideoService
	logx  *logx.Logx
}

func NewVideoHandler(video service.IVideoService, logx *logx.Logx) *VideoHandler {
	return &VideoHandler{
		video: video,
		logx:  logx,
	}
}

func (h *VideoHandler) Start(srv *gin.RouterGroup) {
	srv.GET("/videos/:class_id", h.GetVideoByClassID)
	srv.GET("/videos/detail/:id", h.GetVideoByID)

}

func (h *VideoHandler) GetVideoByClassID(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	idStr := ctx.Params.ByName("class_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err, http.StatusText(http.StatusBadRequest), nil)
		return
	}

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

	videos, err := h.video.GetVideoByClassID(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = videos
}

func (h *VideoHandler) GetVideoByID(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	idStr := ctx.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err, http.StatusText(http.StatusBadRequest), nil)
		return
	}

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

	video, err := h.video.GetVideoByID(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = video
}
