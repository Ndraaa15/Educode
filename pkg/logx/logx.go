package logx

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/Ndraaa15/IQuest/pkg/errsx"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Logx struct {
	Logger *logrus.Logger
}

func NewLogger() (*Logx, error) {
	f, err := os.OpenFile("log/log.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}

	newLogger := logrus.New()
	newLogger.SetLevel(logrus.DebugLevel)
	newLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false,
		TimestampFormat:  "02 Jan 06 - 15:04",
		FullTimestamp:    true,
		CallerPrettyfier: customCallerFormatter,
	})

	newLogger.Out = io.MultiWriter(os.Stdout, f)
	newLogger.SetReportCaller(true)

	return &Logx{
		Logger: newLogger,
	}, nil
}

func customCallerFormatter(f *runtime.Frame) (string, string) {
	s := strings.Split(f.Function, ".")
	funcName := s[len(s)-1]
	fileLine := fmt.Sprintf("[%s:%d]", path.Base(f.File), f.Line)
	return fileLine, funcName + "()"
}

func (l *Logx) RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		l.Logger.WithFields(logrus.Fields{
			"ip":     c.ClientIP(),
			"method": c.Request.Method,
			"uri":    c.Request.URL,
		}).Info("Request")
		c.Next()
	}
}

func (l *Logx) ErrorLogger(err error) {
	var errsx errsx.Errsx
	if errors.As(err, &errsx) {
		log.Println("HI")
		l.Logger.WithFields(logrus.Fields{
			"code":     errsx.Code,
			"location": errsx.Location,
			"error":    errsx.Err,
		}).Error("Process")
	} else {
		l.Logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Process")
	}
}
