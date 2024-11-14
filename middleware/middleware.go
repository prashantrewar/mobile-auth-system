package middleware

import (
    "context"
    "net/http"
    "strings"
    "mobile-auth-system/utils"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        userID, err := utils.ParseJWT(tokenString)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Add user ID to context
        ctx := context.WithValue(r.Context(), "userID", userID)
        r = r.WithContext(ctx)

        // Call the next handler
        next(w, r)
    }
}
