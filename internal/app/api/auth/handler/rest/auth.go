package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/Ndraaa15/IQuest/internal/app/api/auth/service"
	"github.com/Ndraaa15/IQuest/internal/pkg/dto"
	"github.com/Ndraaa15/IQuest/pkg/logx"
	"github.com/Ndraaa15/IQuest/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth service.IAuthService
	logx *logx.Logx
}

func NewAuthHandler(auth service.IAuthService, logx *logx.Logx) *AuthHandler {
	return &AuthHandler{
		auth: auth,
		logx: logx,
	}
}

func (h *AuthHandler) Start(srv *gin.RouterGroup) {
	srv.POST("/auth/signup", h.SignUp)
	srv.POST("/auth/signin", h.SignIn)
}

func (h *AuthHandler) SignUp(ctx *gin.Context) {
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

	var req dto.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errx = err
		code = http.StatusBadRequest
		return
	}

	if err := h.auth.SignUp(c, req); err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	var (
		errx error
		code int = http.StatusOK
		data interface{}
	)

	defer func() {
		if errx != nil {
			h.logx.ErrorLogger(errx)
			response.Error(ctx, code, errx, errx.Error(), nil)
			return
		}
		response.Success(ctx, code, http.StatusText(http.StatusOK), data)
	}()

	var req dto.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errx = err
		code = http.StatusBadRequest
		return
	}

	token, id, err := h.auth.SignIn(c, req)
	if err != nil {
		errx = err
		code = http.StatusInternalServerError
		return
	}

	data = gin.H{
		"token":    token,
		"class_id": id,
	}
}
