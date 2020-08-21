// Handlers file
package main

import (
	"fmt"
	//"github.com/gofiber/template/html"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	Conn *gorm.DB
)

// api
func signup(c *fiber.Ctx) {
	db := Conn

	// Create
	//db.Create(&Product{Code: "L1212", Price: 1000})
	db.AutoMigrate(&User{})

	user := &User{}

	// TODO : hand zero paramse; //  tweet.Title = c.Params("title"); tweet.Body = c.Params("body")

	if err := c.BodyParser(*user); err != nil {
		c.Status(503).JSON(err)
		//fmt.Println(err)
		return
	}

	fmt.Println(user)
	db.Create(&User{Username: "hello", Password: "world"}) //&User{user.Username, user.Password, user.Email, user.Phon, user.Avatarlink})
	c.Send(user)                                           // or c.Send("success")
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

func acount(c *fiber.Ctx) {
	c.Render("acount", nil) // fiber.Map{}
}
func post(c *fiber.Ctx) {
	c.Render("post", nil) // fiber.Map{}
	// TODO apload item
}
