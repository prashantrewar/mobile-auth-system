package services

import (
	"context"
	"fmt"
	"mobile-auth-system/config"
	"mobile-auth-system/utils"
	"strconv"
	"time"
	"math/rand"

)

// OTP expiration and cooldown periods in seconds
const otpExpiry = 300       // 5 minutes
const resendCooldown = 60   // 1 minute cooldown for OTP resend

func SendOTP(mobile string) (string, error) {
    otp := utils.GenerateOTP()
    key := fmt.Sprintf("otp:%s", mobile)
    err := config.RDB.Set(context.Background(), key, otp, 5*time.Minute).Err()
    if err != nil {
        return "", err
    }
    // Send OTP via SMS service (e.g., Twilio)
    return otp, nil
}

func VerifyOTP(mobile, otp string) bool {
    key := fmt.Sprintf("otp:%s", mobile)
    val, err := config.RDB.Get(context.Background(), key).Result()
    if err != nil || val != otp {
        return false
    }
    config.RDB.Del(context.Background(), key)
    return true
}


// GenerateAndStoreOTP generates an OTP, sets an expiration, and applies a resend cooldown
func GenerateAndStoreOTP(mobile string) (string, error) {
	otp := strconv.Itoa(100000 + rand.Intn(900000)) // Generates a 6-digit OTP
	otpKey := fmt.Sprintf("otp:%s", mobile)
	cooldownKey := fmt.Sprintf("cooldown:%s", mobile)

	// Check if resend cooldown is active
	if ttl, err := config.RDB.TTL(context.Background(), cooldownKey).Result(); err == nil && ttl > 0 {
		return "", fmt.Errorf("OTP resend cooldown active, please wait %d seconds", int(ttl.Seconds()))
	}

	// Store OTP with expiration in Redis
	err := config.RDB.Set(context.Background(), otpKey, otp, time.Duration(otpExpiry)*time.Second).Err()
	if err != nil {
		return "", err
	}

	// Set a cooldown for OTP resend requests
	config.RDB.Set(context.Background(), cooldownKey, "cooldown", time.Duration(resendCooldown)*time.Second)

	// Simulate sending OTP via SMS (replace with actual SMS integration)
	fmt.Printf("OTP sent to %s: %s\n", mobile, otp) // This line would be replaced with an SMS API call
	return otp, nil
}