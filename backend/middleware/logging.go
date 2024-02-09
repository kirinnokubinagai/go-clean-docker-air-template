package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// インターフェース
type ILoggingMiddleware interface {
	GetLoggingMiddleware() gin.HandlerFunc
}

type LoggingMiddleware struct{}

// コンストラクター
func NewLoggingMiddleware() ILoggingMiddleware {
	return &LoggingMiddleware{}
}

// ログ用のミドルウェアを返す
// gin.HandlerFunc ログ用のコールバック関数
func (LoggingMiddleware LoggingMiddleware) GetLoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		logrus.WithFields(logrus.Fields{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Info("HTTP REQUEST")
	}
}
