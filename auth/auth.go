package auth

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("supersecretkey")

func GenerateToken() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    return token.SignedString(secretKey)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return secretKey, nil
        })

        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next(w, r)
    }
}