package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/joho/godotenv"

	routers "github.com/noragalvin/go-server/routes"
)

func main() {
	routers := routers.InitRoutes()
	http.Handle("/", routers)

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
