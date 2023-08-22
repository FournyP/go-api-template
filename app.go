package main

import (
	"go-api-template/controllers"
	"go-api-template/settings"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	port          int
	fiberInstance *fiber.App
}

func NewApp(appSettings *settings.AppSettings, helloController *controllers.HelloController) *App {
	app := fiber.New()

	app.Get("/", helloController.GetHello)

	return &App{
		port:          appSettings.Port,
		fiberInstance: app,
	}
}
