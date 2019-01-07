package models

import (
	"fmt"

	_ "github.com/bmizerany/pq"
	"github.com/jinzhu/gorm"
	config "github.com/noragalvin/go-server/app/utils/config"
)

var db *gorm.DB

func init() {
	DBName := config.Get().Database.Name
	DBHost := config.Get().Database.Host
	DBPort := config.Get().Database.Port
	DBUsername := config.Get().Database.Username
	DBPassword := config.Get().Database.Password
	DBString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUsername, DBName, DBPassword)

	conn, err := gorm.Open("postgres", DBString)
	// defer conn.Close()
	if err != nil {
		panic(err)
	}

	db = conn

	db.AutoMigrate(&User{})
}

// OpenDB Connect to database
func OpenDB() *gorm.DB {
	return db
}
