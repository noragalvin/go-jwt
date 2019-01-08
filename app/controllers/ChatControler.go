package controllers

import (
	"net/http"
)

// ChatRoom GET Chat
func ChatRoom(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./app/ws/Home.html")
}
