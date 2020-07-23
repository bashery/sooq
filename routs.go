// routs.go file
package main

import (
	"github.com/gofiber/fiber"
)

// all functions hold in handler.go file
func SetRouts(app *fiber.App) {
	// TODO authentication

	app.Static("/", "./assets")
	app.Get("/home", home)
	app.Get("/login", login)
	app.Get("/sign", sign)
	app.Get("/:test", params)

	app.Post("/body", body)
}
