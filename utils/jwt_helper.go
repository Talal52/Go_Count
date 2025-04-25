package utils

import (
    "time"

    "github.com/golang-jwt/jwt"
)

var secretKey = []byte("secret-key")

func GenerateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString(secretKey)
}