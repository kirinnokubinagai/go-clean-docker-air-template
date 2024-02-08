package server

import (
	"go_app/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Controllers struct {
	dig.In
	User controller.IUserController
}

type IRouter interface {
	setRoutes(router *gin.Engine)
}

type Router struct {
	controllers Controllers
}

func NewRouter(controllers Controllers) IRouter {
	return &Router{controllers}
}

func (router *Router) setRoutes(engine *gin.Engine) {
	engine.GET("/user_list", router.controllers.User.GetUserList)
}
