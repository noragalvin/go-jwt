package routes

import (
	"github.com/gorilla/mux"
	"github.com/noragalvin/go-server/app/controllers"
	"github.com/noragalvin/go-server/app/middleware"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", controllers.UserTest).Methods("GET")
	router.HandleFunc("/api/user", middleware.Authentication(controllers.UserGet)).Methods("GET")
	router.HandleFunc("/api/login", controllers.UserLogin).Methods("POST")

	return router
}
