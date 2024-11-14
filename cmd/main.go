package main

import (
	"log"
	"mobile-auth-system/config"
	"mobile-auth-system/models"
	"mobile-auth-system/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize configuration, DB, and Redis
    config.Initialize()
	config.DB.AutoMigrate(&models.User{})
    router := mux.NewRouter()
    routes.SetupRoutes(router)

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
