// routs.go file
package main

import (
	"github.com/gofiber/fiber"
)

// all functions hold in handler.go file
func SetRouts(app *fiber.App) {
	app.Get("/", home)
	app.Get("/login", login)
	app.Get("/sign", sign)
	app.Get("/:test", params)

	app.Post("/body", body)
}
