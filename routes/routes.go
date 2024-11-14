package routes

import (
	"mobile-auth-system/controllers"
	"mobile-auth-system/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
    router.HandleFunc("/register", controllers.Register).Methods("POST")
    // router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/login/request-otp", controllers.RequestOTP).Methods("POST")
	router.HandleFunc("/login/verify-otp", controllers.VerifyOTP).Methods("POST")
	router.HandleFunc("/login/resend-otp", controllers.ResendOTP).Methods("POST")
	router.HandleFunc("/user", middleware.Authenticate(controllers.GetUser)).Methods("GET")
}
