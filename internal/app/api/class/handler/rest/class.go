package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Ndraaa15/Educode/internal/app/api/class/service"
	"github.com/Ndraaa15/Educode/internal/pkg/dto"
	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"github.com/Ndraaa15/Educode/middleware"
	"github.com/Ndraaa15/Educode/pkg/logx"
	"github.com/Ndraaa15/Educode/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	class service.IClassService
	logx  *logx.Logx
}

func NewClassHandler(course service.IClassService, logx *logx.Logx) *ClassHandler {
	return &ClassHandler{
		class: course,
		logx:  logx,
	}
}

func (h *ClassHandler) Start(srv *gin.RouterGroup) {
	srv.GET("/classes/:id", middleware.ValidateJWTToken("student", "teacher"), h.GetClassByID)
	srv.PATCH("/classes/:id", middleware.ValidateJWTToken("teacher"), h.UpdateClass)
}

func (h *ClassHandler) GetClassByID(ctx *gin.Context) {
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
		code int = http.StatusOK
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

	class, err := h.class.GetClassByID(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = class
}

func (h *ClassHandler) UpdateClass(ctx *gin.Context) {
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
		code int = http.StatusOK
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

	class, err := h.class.GetClassByID(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	var req dto.UpdateClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errx = err
		code = http.StatusBadRequest
		return
	}

	ParseUpdateClassRequest(&req, &class)

	err = h.class.UpdateClass(c, class)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}
}

func ParseUpdateClassRequest(req *dto.UpdateClassRequest, data *entity.Class) {
	if req.Name != "" {
		data.Name = req.Name
	}

	if req.Goal != "" {
		data.Goal = req.Goal
	}
}
