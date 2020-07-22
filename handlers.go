// Handlers file
package main

import (
	"github.com/gofiber/fiber"
)

func params(c *fiber.Ctx) {
	data := c.Params("test")
	c.JSON(data)
}
func body(c *fiber.Ctx) {
	data := c.Body()
	c.JSON(data)
}
func home(c *fiber.Ctx) {
	c.Send("Home page")
}

func login(c *fiber.Ctx) {
	c.Send("login page")

}

func sign(c *fiber.Ctx) {
	c.Send("sign up page")
}
