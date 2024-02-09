package server

import (
	"github.com/gin-gonic/gin"
)

type IServer interface {
	GetEngine() *gin.Engine
}

type Server struct {
	engine *gin.Engine
}

func NewServer(router IRouter, logger ILogger) IServer {
	engine := gin.New()

	router.setRoutes(engine)

	logger.SetLogger(engine)

	return &Server{engine}
}

func (server *Server) GetEngine() *gin.Engine {
	return server.engine
}
