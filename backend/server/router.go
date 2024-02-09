package server

import (
	"go_app/controller"
	"go_app/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// コントローラー一覧
type Controllers struct {
	// パラメーターオブジェクト化してまとめてNewRouterの引数に
	// 参考) https://pkg.go.dev/go.uber.org/dig#section-readme
	dig.In

	// UserControllerのインターフェース
	User controller.IUserController
}

// ミドルウェア一覧
type Middlewares struct {
	dig.In

	Logging middleware.ILoggingMiddleware
}

// インターフェース
type IRouter interface {
	setRoutes(router *gin.Engine)
}

type Router struct {
	// コントローラー一覧
	controllers Controllers

	// ミドルウェア一覧
	middlewares Middlewares
}

// コンストラクター
func NewRouter(controllers Controllers, middlewares Middlewares) IRouter {
	return &Router{controllers, middlewares}
}

// コントローラーとミドルウェアの一覧を取得しながらルーティングの設定
// @return engine *gin.Engine ginを返す
func (router *Router) setRoutes(engine *gin.Engine) {
	engine.Use(router.middlewares.Logging.GetLoggingMiddleware())
	engine.GET("/user_list", router.controllers.User.GetUserList)
}
