package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noragalvin/go-server/app/controllers"
	"github.com/noragalvin/go-server/app/middleware"
	ws "github.com/noragalvin/go-server/app/ws"
)

// InitRoutes initialize routers
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	hub := ws.NewHub()

	go hub.Run()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./app/ws/index.html")
	})

	router.HandleFunc("/chat", controllers.ChatRoom).Methods("GET")

	router.Handle("/ws", hub.Handler())

	router.HandleFunc("/api/user/{id}", middleware.Authentication(controllers.UserGet)).Methods("GET")
	router.HandleFunc("/api/login", controllers.UserLogin).Methods("POST")
	router.HandleFunc("/api/register", controllers.UserRegister).Methods("POST")

	return router
}
