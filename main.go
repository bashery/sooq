package main

import (
	//"fmt"
	//"./routs"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", Hello)

	app.Listen(9000)
}
