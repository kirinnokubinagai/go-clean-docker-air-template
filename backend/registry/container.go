package registry

import (
	"go_app/controller"
	"go_app/database"
	"go_app/middleware"
	"go_app/repository"
	"go_app/server"
	"go_app/service"

	"go.uber.org/dig"
)

type IContainer interface {
	// DIコンテナの作成
	BuildContainer() (*dig.Container, error)
}

type Container struct{}

func NewContainer() IContainer {
	return &Container{}
}

// DIコンテナの作成
// @return *dig.Container DIコンテナ
// @return error エラー
func (c *Container) BuildContainer() (*dig.Container, error) {
	container := dig.New()

	providers := []interface{}{
		database.NewDbConnection,
		controller.NewUserController,
		service.NewUserService,
		repository.NewUserRepository,
		middleware.NewLoggingMiddleware,
		server.NewLogger,
		server.NewRouter,
		server.NewServer,
	}

	for _, provider := range providers {
		if err := container.Provide(provider); err != nil {
			return nil, err
		}
	}

	return container, nil
}
