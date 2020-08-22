// Handlers file
package main

import (
	"fmt"
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
	var user User
	//db.AutoMigrate(&User{})

	if err := c.BodyParser(&user); err != nil {
		c.Status(503).Send(err)
		//fmt.Println(err)
		return
	}

	fmt.Printf("user name is : %T\n", user.Username)
	fmt.Printf("user : %T\n", user)
	fmt.Printf("&user : %T\n", &user)
	fmt.Println(&user)
	user = User{Username: user.Username, Password: user.Password, Email: user.Email, Phon: user.Email, Avatarlink: user.Avatarlink}
	db.Create(&user) //User{user.Username, user.Password, user.Email, user.Phon, user.Avatarlink})
	c.Send(&user)    // or c.Send("success")

	// Create
	//db.Create(&Product{Code: "L1212", Price: 1000})
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
