package controllers

import "net/http"

func ChatRoom(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./app/ws/Home.html")
}
