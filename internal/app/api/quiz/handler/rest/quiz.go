package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Ndraaa15/Educode/internal/app/api/quiz/service"
	"github.com/Ndraaa15/Educode/internal/pkg/dto"
	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"github.com/Ndraaa15/Educode/middleware"
	"github.com/Ndraaa15/Educode/pkg/logx"
	"github.com/Ndraaa15/Educode/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	course service.IQuizService
	logx   *logx.Logx
}

func NewQuizHandler(course service.IQuizService, logx *logx.Logx) *QuizHandler {
	return &QuizHandler{
		course: course,
		logx:   logx,
	}
}

func (h *QuizHandler) Start(srv *gin.RouterGroup) {
	srv.GET("/quizzes/:class_id", middleware.ValidateJWTToken("student", "teacher"), h.GetQuizByClassID)
	srv.GET("/quizzes/detail/:id", middleware.ValidateJWTToken("student", "teacher"), h.GetQuizByID)
	srv.POST("/quizzes/attempt", middleware.ValidateJWTToken("student"), h.AttemptQuiz)
	srv.GET("/quizzes/result/:id", middleware.ValidateJWTToken("teacher"), h.GetQuizResult)
}

func (h *QuizHandler) GetQuizByClassID(ctx *gin.Context) {
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

	quizzes, err := h.course.GetQuizByClassID(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = quizzes
}

func (h *QuizHandler) GetQuizByID(ctx *gin.Context) {
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

	quiz, err := h.course.GetQuizByID(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = quiz
}

func (h *QuizHandler) AttemptQuiz(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	user, ok := ctx.Get("user")
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, nil, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	u, ok := user.(*entity.JWTClaims)
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, nil, http.StatusText(http.StatusUnauthorized), nil)
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
		response.Success(ctx, code, http.StatusText(http.StatusOK), data)
	}()

	var req dto.AttemptQuizRequest
	if err := ctx.BindJSON(&req); err != nil {
		errx = err
		code = http.StatusBadRequest
		return
	}

	result, err := h.course.AttemptQuiz(c, req, u.ID)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = result
}

func (h *QuizHandler) GetQuizResult(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
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
		response.Success(ctx, code, http.StatusText(http.StatusOK), data)
	}()

	result, err := h.course.GetQuizResult(c, id)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = result
}
