package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func getServerStatus(c *gin.Context) {
	// Получение статуса VPN-сервера
	c.JSON(200, gin.H{"status": "VPN server is running"})
}

func getServerStats(c *gin.Context) {
	// Статистика работы VPN-сервера
	c.JSON(200, gin.H{
		"uptime":    "72 hours",
		"load":      "15%",
		"connections": "256 active connections",
	})
}

func main() {
	r := gin.Default()

	r.GET("/status", getServerStatus)
	r.GET("/stats", getServerStats)

	log.Fatal(r.Run(":8081"))
}
