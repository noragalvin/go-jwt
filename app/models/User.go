package models

import (
	"log"
)

type User struct {
	ID        string `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expied_at`
}

func NewUser() *User {
	return &(User{})
}

func (u *User) UserShow(id uint) User {
	var user User
	db := OpenDB()
	db.Where("id = ?", id).First(&user)
	return user
}

func (u *User) Test() {
	log.Println("Hi")
}
