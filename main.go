package main

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Welcome in begger sooq in the world")
	})

	app.Listen(3000)
}
