package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	"github.com/jinzhu/gorm"
	"log"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, err := gorm.Open("sqlite3", "../test.db")
	//db, err := gorm.Open("sqlite3", "../test.db")

	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	engine := html.New("templates", ".html")
	engine.Reload(true) //engine.Debug(true)

	app := fiber.New(&fiber.Settings{Views: engine})

	SetRouts(app)
	app.Listen(9000)

}
