// cmd/vpn-server/auth.go
package main

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// Секретный ключ для подписи JWT
var jwtKey = []byte("my_secret_key")

// Структура для хранения данных пользователя
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Структура для хранения токена
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Функция для создания нового токена
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Функция для проверки токена
func ValidateJWT(tokenStr string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	return claims.Username, nil
}
