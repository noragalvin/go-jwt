package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/noragalvin/go-server/app/utils/auth"

	"github.com/asaskevich/govalidator"
	helpers "github.com/noragalvin/go-server/app/helpers"
	models "github.com/noragalvin/go-server/app/models"
	view "github.com/noragalvin/go-server/app/utils/view"
)

// type Test struct {
// 	A string `json:"a"`
// 	B string `json:"b"`
// }

// UserRegister signup user
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var userReq models.UserNoPwd

	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		view.Respond(w, view.Message(false, err.Error()))
		return
	}

	//Validate
	if _, err := govalidator.ValidateStruct(userReq); err != nil {
		view.Respond(w, view.Message(false, err.Error()))
		return
	}

	db := models.OpenDB()
	var check models.User

	if err := db.Where("email = ?", userReq.Email).Or("username = ?", userReq.Username).First(&check).Error; gorm.IsRecordNotFoundError(err) {
		view.Respond(w, view.Message(false, err.Error()))
		return
	}

	//Create user
	var user models.User
	user.Username = userReq.Username
	user.Email = userReq.Email
	user.Password = helpers.HashAndSalt(userReq.Password)
	user.Fullname = userReq.Fullname
	db.Create(&user)
	// tokenString, expriredAt := auth.CreateToken(user.ID, user.Email, user.Fullname)

	data := view.Message(true, "success")
	// data["token"] = tokenString
	// data["expired"] = expriredAt
	view.Respond(w, data)
	return

}

// UserLogin login
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var user models.UserNoPwd
	// Decode request data to user variable
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		return
	}

	db := models.OpenDB()
	var userDB models.User
	// Find exist record in database
	if err := db.Where("email = ?", user.Email).First(&userDB).Error; gorm.IsRecordNotFoundError(err) {
		view.Respond(w, view.Message(false, err.Error()))
		return
	}

	//Compare request password and exist password
	if checkPwd := helpers.ComparePassword(userDB.Password, user.Password); !checkPwd {
		view.Respond(w, view.Message(false, "Incorrect password"))
		return
	}

	tokenString, expiresAt := auth.CreateToken(user.ID, user.Email, user.Fullname)

	data := view.Message(true, "success")
	data["token"] = tokenString
	data["expires"] = expiresAt
	view.Respond(w, data)
	return
}

// UserGet Get an user by ID
func UserGet(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var id = mux.Vars(r)["id"]
	// log.Println(id)
	db := models.OpenDB()
	db.Where("id = ?", id).First(&user)
	data := view.Message(true, "success")
	data["user"] = user
	view.Respond(w, data)
}
