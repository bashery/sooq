package main

//import "github.com/jinzhu/gorm"

// type adAdminfor register admin inrormation
type Admin struct{}

// User type model for register user information
type User struct {
	//gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Phon       string `json:"phon"`
	Avatarlink string `json:"avatarlink"`
}
