package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
}

type Server struct {
	*gin.Engine
}

func NewServer(router IRouter) *Server {
	engine := gin.New()

	engine.Use()
	router.setRoutes(engine)

	return &Server{engine}
}
