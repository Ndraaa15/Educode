package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Ndraaa15/IQuest/internal/app/api/course/service"
	"github.com/Ndraaa15/IQuest/internal/pkg/dto"
	"github.com/Ndraaa15/IQuest/middleware"
	"github.com/Ndraaa15/IQuest/pkg/logx"
	"github.com/Ndraaa15/IQuest/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	course service.ICourseService
	logx   *logx.Logx
}

func NewCourseHandler(course service.ICourseService, logx *logx.Logx) *CourseHandler {
	return &CourseHandler{
		course: course,
		logx:   logx,
	}
}

func (h *CourseHandler) Start(srv *gin.RouterGroup) {
	srv.POST("/courses", middleware.ValidateJWTToken("teacher"), h.CreateCourse)
	srv.GET("/courses/:class_id", middleware.ValidateJWTToken("student", "teacher"), h.GetCourseByClassID)
	srv.GET("/courses/detail/:id", middleware.ValidateJWTToken("student", "teacher"), h.GetCourseByID)
}

func (h *CourseHandler) CreateCourse(ctx *gin.Context) {
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

	var req dto.CreateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errx = err
		code = http.StatusBadRequest
		return
	}

	if err := h.course.CreateCourse(c, req); err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}
}

func (h *CourseHandler) GetCourseByClassID(ctx *gin.Context) {
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

	class, err := h.course.GetCourseByClassID(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = class
}

func (h *CourseHandler) GetCourseByID(ctx *gin.Context) {
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

	class, err := h.course.GetCourseByID(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = class
}
