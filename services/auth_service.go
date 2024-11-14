package services

import (
	"errors"
	"mobile-auth-system/config"
	"mobile-auth-system/models"
)

var ErrUserNotFound = errors.New("user not found")

func RegisterUser(mobile, fingerprint string) (*models.User, error) {
    user := &models.User{MobileNumber: mobile, DeviceFingerprint: fingerprint}
    if err := config.DB.Create(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

// VerifyFingerprint checks if the user’s fingerprint matches the stored fingerprint
func VerifyFingerprint(mobile, fingerprint string) (bool, error) {
    user, err := GetUserByMobile(mobile)
    if err != nil {
        return false, errors.New("user not found")
    }
    if user.DeviceFingerprint != fingerprint {
        // You could log this, send a notification, or require extra verification here
        return false, errors.New("device fingerprint mismatch: new device detected")
    }
    return true, nil
}

func GetUserByMobile(mobile string) (*models.User, error) {
    var user models.User
    if err := config.DB.Where("mobile_number = ?", mobile).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// GetUserByID retrieves a user by their ID from the database.
func GetUserByID(userID uint) (*models.User, error) {
    var user models.User
    result := config.DB.First(&user, userID)
    if result.Error != nil {
        return nil, ErrUserNotFound
    }
    return &user, nil
}