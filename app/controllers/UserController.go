package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	helpers "github.com/noragalvin/go-server/app/helpers"
	models "github.com/noragalvin/go-server/app/models"
)

const username string = "minhnora98"
const hashPwd string = "$2a$04$02FxQYX4.ghQ1RDZkbIQNO9JAD1uWs76jx1YoOakvI.7ENqL1XRc2" // "123456"

// type Test struct {
// 	A string `json:"a"`
// 	B string `json:"b"`
// }

func UserRegister(w http.ResponseWriter, r *http.Request) {
	return
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}
	var check bool = helpers.ComparePassword(hashPwd, "123456")
	if username != user.UserName && check == false {
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.UserName,
		"password": user.Password,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	user.Token = tokenString

	json.NewEncoder(w).Encode(user)
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	var user = models.UserShow()
	json.NewEncoder(w).Encode(user)
	// json.NewEncoder(w).Encode(Test{A: "1", B: "2"})
}

func UserTest(w http.ResponseWriter, r *http.Request) {

	// var pwd string = helpers.HashPassword("123456")
	var check bool = helpers.ComparePassword(hashPwd, "123456")
	log.Println(check)
}
