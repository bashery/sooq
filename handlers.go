// Handlers file
package main

import (
	"fmt"
	"github.com/gofiber/fiber"
)

type data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func params(c *fiber.Ctx) {
	//data := c.Params("test")
	//c.JSON(data)
}
func body(c *fiber.Ctx) {
	d := &data{}
	if err := c.BodyParser(d); err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println(d)
	//c.JSON(d)
	c.Send(d)
}
func home(c *fiber.Ctx) {
	_ = c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
	//c.Send("Home page")
}

func login(c *fiber.Ctx) {
	c.Send("login page")

}

func sign(c *fiber.Ctx) {
	c.Send("sign up page")
}
