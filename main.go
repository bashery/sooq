package main

import (
	//"fmt"
	//"./routs"
	"github.com/gofiber/fiber"
)

func Hello(c *fiber.Ctx) {
	c.Send("welcom in beeger sooq in the university")
}
func main() {
	app := fiber.New()

	app.Get("/", Hello)

	app.Listen(9000)
}
