package main

import (
	"fiber-orchestrator/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/", handlers.RegisterSportmen)

	app.Listen(":80")
}
