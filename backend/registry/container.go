package registry

import (
	"go_app/controller"
	"go_app/database"
	"go_app/repository"
	"go_app/server"
	"go_app/service"

	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	providers := []interface{}{
		database.BootDatabase,
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
