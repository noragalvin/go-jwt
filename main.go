package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/joho/godotenv"
	"github.com/noragalvin/go-server/app/utils/config"
	routers "github.com/noragalvin/go-server/routes"
)

func main() {
	routers := routers.InitRoutes()
	http.Handle("/api/", routers)
	config.InitConfig()

	fmt.Println("Server is running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
