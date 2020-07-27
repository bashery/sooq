package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("./assets", ".html")

	app := fiber.New(&fiber.Settings{
		Views: engine,
	})
	SetRouts(app)
	app.Listen(9000)

}
