// cmd/vpn-server/main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Мокированный пользователь для демонстрации
var mockUser = map[string]string{
	"admin": "password123", // username: admin, password: password123
}

// Функция для входа (генерация токена)
func login(c *gin.Context) {
	var credentials Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	// Проверка пользователя
	if password, ok := mockUser[credentials.Username]; ok && password == credentials.Password {
		token, err := GenerateJWT(credentials.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// Мидлвар для проверки авторизации
func authorizeJWT(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	username, err := ValidateJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("username", username)
}

func startVPNServer(c *gin.Context) {
	// Механизм запуска VPN
	username := c.MustGet("username").(string)
	fmt.Printf("User %s is starting VPN server\n", username)
	c.JSON(http.StatusOK, gin.H{"status": "VPN server started"})
}

func main() {
	r := gin.Default()

	// Роут для входа
	r.POST("/login", login)

	// Применяем middleware для аутентификации
	r.Use(authorizeJWT)

	r.GET("/start", startVPNServer)

	r.Run(":8080")
}
