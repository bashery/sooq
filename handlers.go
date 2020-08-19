// Handlers file
package main

import (
	"fmt"
	//"github.com/gofiber/template/html"
	"github.com/gofiber/fiber"
)

type data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func home(c *fiber.Ctx) {
	err := c.Render("index", fiber.Map{"Title": "Hello,World!"})
	if err != nil {
		fmt.Println(err)
	}
	c.Render("home", fiber.Map{"Tilte": "this is a renderd val"})
}

func sign(c *fiber.Ctx) {
	c.Render("sign", nil) //fiber.Map{})
}

func login(c *fiber.Ctx) {
	c.Render("login", nil) //fiber.Map{})
}

func logout(c *fiber.Ctx) {
	// TODO handle cockeis and log out client
	c.Render("index", nil) //fiber.Map{})
}

func stores(c *fiber.Ctx) {
	c.Render("stores", nil) //fiber.Map{})
}

func mystore(c *fiber.Ctx) {
	c.Render("mystore", nil) // fiber.Map{}
}

func profile(c *fiber.Ctx) {
	c.Render("profile", nil) // fiber.Map{}
}
func post(c *fiber.Ctx) {
	c.Render("post", nil) // fiber.Map{}
	// TODO apload item
}

// api
func body(c *fiber.Ctx) {
	d := &data{}
	if err := c.BodyParser(d); err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println(d)
	//c.JSON(d)
	c.Send(d)
}
