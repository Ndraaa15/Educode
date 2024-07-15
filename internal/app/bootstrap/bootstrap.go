package bootstrap

import (
	"fmt"
	"net/http"
	"os"

	authhandler "github.com/Ndraaa15/Educode/internal/app/api/auth/handler/rest"
	authrepository "github.com/Ndraaa15/Educode/internal/app/api/auth/repository"
	authservice "github.com/Ndraaa15/Educode/internal/app/api/auth/service"
	classhandler "github.com/Ndraaa15/Educode/internal/app/api/class/handler/rest"
	classrepository "github.com/Ndraaa15/Educode/internal/app/api/class/repository"
	classservice "github.com/Ndraaa15/Educode/internal/app/api/class/service"
	coursehandler "github.com/Ndraaa15/Educode/internal/app/api/course/handler/rest"
	courserepository "github.com/Ndraaa15/Educode/internal/app/api/course/repository"
	courseservice "github.com/Ndraaa15/Educode/internal/app/api/course/service"
	quizhandler "github.com/Ndraaa15/Educode/internal/app/api/quiz/handler/rest"
	quizrepository "github.com/Ndraaa15/Educode/internal/app/api/quiz/repository"
	quizservice "github.com/Ndraaa15/Educode/internal/app/api/quiz/service"
	uploadhandler "github.com/Ndraaa15/Educode/internal/app/api/upload/handler/rest"
	videohandler "github.com/Ndraaa15/Educode/internal/app/api/video/handler/rest"
	videorepository "github.com/Ndraaa15/Educode/internal/app/api/video/repository"
	videoservice "github.com/Ndraaa15/Educode/internal/app/api/video/service"

	"github.com/Ndraaa15/Educode/middleware"
	"github.com/Ndraaa15/Educode/pkg/cloudinary"
	"github.com/Ndraaa15/Educode/pkg/db/postgresql"
	"github.com/Ndraaa15/Educode/pkg/logx"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Bootstrap struct {
	srv     *gin.Engine
	db      *gorm.DB
	logx    *logx.Logx
	cld     cloudinary.CloudinaryItf
	handler []handler
}

type handler interface {
	Start(srv *gin.RouterGroup)
}

func NewBootstrap() (*Bootstrap, error) {
	logx, err := logx.NewLogger()
	if err != nil {
		return nil, err
	}

	db, err := postgresql.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	if err := postgresql.Drop(db); err != nil {
		return nil, err
	}

	if err := postgresql.Migrate(db); err != nil {
		return nil, err
	}

	if err = postgresql.MigrateSeed(db); err != nil {
		return nil, err
	}

	cld, err := cloudinary.NewCloudinary()
	if err != nil {
		return nil, err
	}

	bootstrap := &Bootstrap{
		srv: gin.Default(gin.OptionFunc(func(e *gin.Engine) {
			e.Use(gin.Recovery())
		})),
		db:   db,
		logx: logx,
		cld:  cld,
	}

	return bootstrap, nil
}

func (b *Bootstrap) RegisterHandler() {
	authRepository := authrepository.NewAuthRepository(b.db)
	authSvc := authservice.NewAuthService(authRepository)
	authHandler := authhandler.NewAuthHandler(authSvc, b.logx)

	classRepository := classrepository.NewClassRepository(b.db)
	classSvc := classservice.NewClassService(classRepository)
	classHandler := classhandler.NewClassHandler(classSvc, b.logx)

	courseRepository := courserepository.NewCourseRepository(b.db)
	courseSvc := courseservice.NewCourseService(courseRepository)
	courseHandler := coursehandler.NewCourseHandler(courseSvc, b.logx)

	quizRepository := quizrepository.NewQuizRepository(b.db)
	quizSvc := quizservice.NewQuizService(quizRepository)
	quizHandler := quizhandler.NewQuizHandler(quizSvc, b.logx)

	videoRepository := videorepository.NewVideoRepository(b.db)
	videoSvc := videoservice.NewVideoService(videoRepository)
	videoHandler := videohandler.NewVideoHandler(videoSvc, b.logx)

	uploadHandler := uploadhandler.NewUploadHandler(b.cld, b.logx)

	b.CheckHealth()
	b.handler = append(b.handler, authHandler, classHandler, uploadHandler, courseHandler, quizHandler, videoHandler)
}

func (b *Bootstrap) Run() error {
	b.srv.Use(middleware.Cors())
	b.srv.Use(b.logx.RequestLogger())
	srv := b.srv.Group("/api/v1")

	for _, h := range b.handler {
		h.Start(srv)
	}

	addr := os.Getenv("ADDR")
	port := os.Getenv("PORT")

	if err := b.srv.Run(fmt.Sprintf("%s:%s", addr, port)); err != nil {
		return err
	}

	return nil
}

func (h *Bootstrap) CheckHealth() {
	h.srv.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "🚗💨Beep Beep Your Server is Healthy!",
		})
	})
}
