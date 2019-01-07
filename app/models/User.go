package models

import (
	"time"
)

type User struct {
	Model
	Fullname  string     `json:"fullname" gorm:"column:fullname;type:varchar(30)"`
	Username  string     `json:"username" gorm:"column:username;type:varchar(50); not null;unique"`
	Email     string     `json:"email" gorm:"column:email;type:varchar(50);not null;unique"`
	Password  string     `json:"-" gorm:"column:password;type:varchar(100); not null"`
	Token     string     `json:"token" gorm:"column:token;type:varchar(100)"`
	ExpiredAt *time.Time `json:"expied_at" gorm:"column:exprired_at"`
}

type UserNoPwd struct {
	Model
	Fullname string `json:"fullname" valid:"required~Fullname is required,runelength(2|30)~Name must be from 2 to 30 characters"`
	Username string `json:"username" valid:"required~Username is required, runelength(6|15)~Username must be from 6 to 15 characters"`
	Email    string `json:"email" valid:"required~Email is required,email~Email does not match"`
	Password string `json:"password" valid:"required~Password is required, runelength(6|15)~Password must be from 6 to 15 characters"`
}
