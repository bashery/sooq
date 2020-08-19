package main

// routs.go file

import (
	"github.com/gofiber/fiber"
)

// all functions hold in handler.go file
func SetRouts(app *fiber.App) {
	// TODO authentication

	app.Static("/", "assets") //no static in sam path witn template
	app.Get("/", home)
	app.Get("/mystore", mystore)
	app.Get("/profile", profile)
	app.Get("/stores", stores)
	app.Get("/login", login)
	app.Get("/logout", logout)
	app.Get("/post", post)

	app.Get("/sign", sign)

	// api
	app.Post("/api", body)
}
