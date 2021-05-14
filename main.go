package main

import (
	"github.com/demogoo/diggo/internal/config"
	"github.com/demogoo/diggo/internal/handlers"
	"github.com/demogoo/diggo/internal/server"
	"github.com/demogoo/diggo/internal/services"
	"github.com/demogoo/diggo/internal/storage"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// Providing config
	container.Provide(config.NewConfig)

	// Providing storages connection
	container.Provide(storage.NewDBConn)
	container.Provide(storage.NewCacheConn)
	container.Provide(storage.NewMongoDBConn)

	// Providing services
	container.Provide(services.NewItemService)
	container.Provide(services.NewMemberService)
	container.Provide(services.NewOrderService)

	// Providing handlers
	container.Provide(handlers.NewItemHandler)
	container.Provide(handlers.NewMemberHandler)
	container.Provide(handlers.NewOrderHandler)

	// Providing server
	container.Provide(server.NewServer)

	return container
}

func main() {
	container := BuildContainer()
	err := container.Invoke(func(s *server.Server) {
		s.SetupRoutes()
		s.Engine.Run(s.Port)
	})
	if err != nil {
		panic(err)
	}
}
