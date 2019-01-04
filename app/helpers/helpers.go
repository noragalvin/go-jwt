package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	byteString := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(byteString, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func ComparePassword(hashPassword string, plainPass string) bool {
	byteString := []byte(plainPass)
	byteHash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, byteString)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
