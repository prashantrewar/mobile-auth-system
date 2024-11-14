package controllers

import (
	"encoding/json"
	"fmt"
	"mobile-auth-system/services"
	"mobile-auth-system/utils"
	"net/http"
)

type RegisterRequest struct {
    Mobile      string `json:"mobile"`
    Fingerprint string `json:"fingerprint"`
}

type LoginRequest struct {
    Mobile string `json:"mobile"`
    OTP    string `json:"otp"`
	Fingerprint string `json:"fingerprint"`
}

type OTPRequest struct {
    Mobile string `json:"mobile"`
}

type VerifyOTPRequest struct {
    Mobile string `json:"mobile"`
    OTP    string `json:"otp"`
}


type ResendOTPRequest struct {
    Mobile string `json:"mobile"`
}


func Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    json.NewDecoder(r.Body).Decode(&req)
    user, err := services.RegisterUser(req.Mobile, req.Fingerprint)
    if err != nil {
        http.Error(w, "Registration failed", http.StatusBadRequest)
        return
    }
    json.NewEncoder(w).Encode(user)
}


func RequestOTP(w http.ResponseWriter, r *http.Request) {
    var req OTPRequest
    // Decode the incoming request JSON body
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Generate OTP and store it in Redis
    otp, err := services.GenerateAndStoreOTP(req.Mobile)
    if err != nil {
        http.Error(w, "Failed to generate OTP", http.StatusInternalServerError)
        return
    }

    // Log the OTP for debugging (optional)
    // You can remove this in production
    fmt.Println("Generated OTP:", otp)

    // Send a response indicating that the OTP has been sent
    response := map[string]string{
        "message": "OTP sent",
        "otp":     otp,  // Optionally include the OTP in the response
    }
    json.NewEncoder(w).Encode(response)
}


func VerifyOTP(w http.ResponseWriter, r *http.Request) {
    var req VerifyOTPRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Verify the OTP
    isValid := services.VerifyOTP(req.Mobile, req.OTP)
    if !isValid {
        http.Error(w, "Invalid or expired OTP", http.StatusUnauthorized)
        return
    }

    // Fetch user information and generate JWT token
    user, _ := services.GetUserByMobile(req.Mobile)
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    // Send response with JWT token
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}


func ResendOTP(w http.ResponseWriter, r *http.Request) {
    var req ResendOTPRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Generate and store a new OTP
    otp, err := services.GenerateAndStoreOTP(req.Mobile)
    if err != nil {
        http.Error(w, "Failed to resend OTP", http.StatusInternalServerError)
        return
    }

    // Optionally log the OTP for debugging
    fmt.Println("Resent OTP:", otp)

    // Send a response indicating that the OTP has been resent
    response := map[string]string{
        "message": "OTP resent",
		"otp":     otp,
    }
    json.NewEncoder(w).Encode(response)
}


func Login(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    json.NewDecoder(r.Body).Decode(&req)

    if !services.VerifyOTP(req.Mobile, req.OTP) {
        http.Error(w, "Invalid OTP", http.StatusUnauthorized)
        return
    }

    // Verify device fingerprint
    if matched, err := services.VerifyFingerprint(req.Mobile, req.Fingerprint); !matched {
        http.Error(w, "New device detected: " + err.Error(), http.StatusUnauthorized)
        return
    }

    user, _ := services.GetUserByMobile(req.Mobile)
    token, _ := utils.GenerateJWT(user.ID)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}



func GetUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from the context (added by the authentication middleware)
    userID, err := utils.GetUserIDFromContext(r.Context())
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Fetch user details using the userID
    user, err := services.GetUserByID(userID)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Return user details as JSON
    json.NewEncoder(w).Encode(user)
}