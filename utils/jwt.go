package utils

import (
	"context"
	"errors"
	// "os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_jwt_secret_key")

func GenerateJWT(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(jwtSecret)
}

// ParseJWT parses and validates a JWT token, returning the userID if valid.
func ParseJWT(tokenString string) (uint, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtSecret, nil
    })
    
    if err != nil {
        return 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        exp := int64(claims["exp"].(float64))
        if time.Now().Unix() > exp {
            return 0, errors.New("token expired")
        }
        userID := uint(claims["user_id"].(float64))
        return userID, nil
    }

    return 0, errors.New("invalid token")
}

var ErrUserIDNotFound = errors.New("user ID not found in context")

// GetUserIDFromContext retrieves the user ID from the context.
func GetUserIDFromContext(ctx context.Context) (uint, error) {
    userID, ok := ctx.Value("userID").(uint)
    if !ok {
        return 0, ErrUserIDNotFound
    }
    return userID, nil
}