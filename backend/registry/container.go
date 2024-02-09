package registry

import (
	"go_app/controller"
	"go_app/database"
	"go_app/repository"
	"go_app/server"
	"go_app/service"

	"go.uber.org/dig"
)

type IContainer interface {
	BuildContainer() (*dig.Container, error)
}

type Container struct{}

func NewContainer() IContainer {
	return &Container{}
}

// BuildContainer メソッドを実装
func (c *Container) BuildContainer() (*dig.Container, error) {
	container := dig.New()

	providers := []interface{}{
		database.NewDbConnection,
		controller.NewUserController,
		service.NewUserService,
		repository.NewUserRepository,
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
