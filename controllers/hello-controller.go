package controllers

import (
	"go-api-template/ent"

	"github.com/gofiber/fiber/v2"
)

type HelloController struct {
	client *ent.Client
}

func NewHelloController(client *ent.Client) *HelloController { // Get services as you want thanks to wire
	return &HelloController{
		client: client,
	}
}

func (controller *HelloController) GetHello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
