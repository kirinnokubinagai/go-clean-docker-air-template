package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	SetLogger(server *gin.Engine)
}

type Logger struct{}

// コンストラクター
func NewLogger() ILogger {
	return &Logger{}
}

// logrusの設定を付与
func (logger Logger) SetLogger(server *gin.Engine) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}
