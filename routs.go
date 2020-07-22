package main

import (
	"github.com/gofiber/fiber"
)

func Hello(c *fiber.Ctx) {
	c.Send("hello for all")
}
