package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/noragalvin/go-server/app/utils/config"
)

var db *gorm.DB

func InitDB() {
	DBName := config.Get().Database.DBName
	DBHost := config.Get().Database.DBHost
	DBPort := config.Get().Database.DBPort
	DBUsername := config.Get().Database.DBUsername
	DBPassword := config.Get().Database.DBPassword

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUsername, DBName, DBPassword)

	conn, err := gorm.Open("postgres", dbString)
	defer conn.Close()
	if err != nil {
		log.Println(err)
	}
	db = conn

	db.Debug().AutoMigrate(&User{})
}

func OpenDB() *gorm.DB {
	return db
}
