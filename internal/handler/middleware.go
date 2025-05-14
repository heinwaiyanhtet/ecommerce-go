// internal/handlers/middleware.go
package handlers

import (
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(secret []byte) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            auth := r.Header.Get("Authorization")
            if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
                http.Error(w, "missing or invalid auth header", http.StatusUnauthorized)
                return
            }
            tokenStr := strings.TrimPrefix(auth, "Bearer ")
            token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
                return secret, nil
            })
            if err != nil || !token.Valid {
                http.Error(w, "invalid token", http.StatusUnauthorized)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
