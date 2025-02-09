package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	if _, err := bot.Send(msg); err != nil {
		log.Println("Error sending message:", err)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("YOUR_BOT_API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/send-update", func(c *gin.Context) {
		sendMessage(bot, 123456789, "VPN server status update")
		c.JSON(200, gin.H{"message": "Update sent"})
	})

	log.Fatal(r.Run(":8082"))
}
