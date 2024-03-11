package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	type Sub struct {
		Plan  string `json:"plan"`
		Price int    `json:"price"`
	}

	var SubJson []Sub
	app := fiber.New()

	app.Use(logger.New())

	sub1 := Sub{Plan: "basico", Price: 0}
	sub2 := Sub{Plan: "intermedio", Price: 19}
	sub3 := Sub{Plan: "premium", Price: 39}

	SubJson = append(SubJson, sub1, sub2, sub3)

	app.Get("/sub/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("ping")
	})

	app.Get("/sub", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(SubJson)
	})

	app.Listen(":80")

}
