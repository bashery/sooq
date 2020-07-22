package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	SetRouts(app)
	app.Listen(9000)
}
