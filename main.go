package main

import (
	"github.com/Guidotss/ia-saas-backend.git/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configuration := config.New()
	database := config.NewDatabase(configuration)

	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Hello, World!")
	})

	app.Listen(":" + configuration.Get("PORT"))
}
