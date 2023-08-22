//go:build wireinject
// +build wireinject

package main

import (
	"go-api-template/controllers"
	"go-api-template/settings"

	"github.com/google/wire"
)

func InitializeApp() *App {
	wire.Build(
		NewApp,
		NewDatabase,
		settings.NewAppSettings,
		settings.NewDatabaseSettings,
		controllers.NewHelloController,
	)

	return &App{}
}
