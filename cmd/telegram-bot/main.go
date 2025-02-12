package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "github.com/sirupsen/logrus"
)

var (
    botToken   = os.Getenv("TELEGRAM_BOT_TOKEN")
    adminChat  = os.Getenv("ADMIN_CHAT_ID") // ID чата администратора
    logger     *logrus.Logger
)

func init() {
    // Настройка логгера
    logger = logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})
    logger.SetLevel(logrus.InfoLevel)

    // Проверка обязательных переменных окружения
    if botToken == "" || adminChat == "" {
        logger.Fatalf("Missing required environment variables: TELEGRAM_BOT_TOKEN or ADMIN_CHAT_ID")
    }
}

// sendMessage отправляет текстовое сообщение в указанный чат
func sendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) error {
    msg := tgbotapi.NewMessage(chatID, message)
    _, err := bot.Send(msg)
    if err != nil {
        logger.Errorf("Error sending message to chat %d: %v", chatID, err)
        return err
    }
    logger.Infof("Message sent to chat %d: %s", chatID, message)
    return nil
}

// sendPhoto отправляет изображение в указанный чат
func sendPhoto(bot *tgbotapi.BotAPI, chatID int64, photoURL string) error {
    msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(photoURL))
    _, err := bot.Send(msg)
    if err != nil {
        logger.Errorf("Error sending photo to chat %d: %v", chatID, err)
        return err
    }
    logger.Infof("Photo sent to chat %d: %s", chatID, photoURL)
    return nil
}

// handleBotCommands обрабатывает команды от пользователей
func handleBotCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    if update.Message == nil || !update.Message.IsCommand() {
        return
    }

    command := update.Message.Command()
    chatID := update.Message.Chat.ID

    switch command {
    case "start":
        sendMessage(bot, chatID, "Welcome! Use /status to check the server status.")
    case "status":
        sendMessage(bot, chatID, "VPN server is running.")
    default:
        sendMessage(bot, chatID, "Unknown command.")
    }
}

// startBot запускает Telegram-бота и обрабатывает входящие сообщения
func startBot(bot *tgbotapi.BotAPI) {
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)
    if err != nil {
        logger.Fatalf("Failed to get updates channel: %v", err)
    }

    for update := range updates {
        go handleBotCommands(bot, update)
    }
}

func main() {
    // Инициализация Telegram-бота
    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        logger.Fatalf("Failed to initialize Telegram bot: %v", err)
    }
    logger.Infof("Authorized on account %s", bot.Self.UserName)

    // Запуск бота в отдельной горутине
    go startBot(bot)

    // Инициализация Gin
    r := gin.Default()

    // Маршрут для отправки уведомлений
    r.GET("/send-update", func(c *gin.Context) {
        message := c.Query("message")
        if message == "" {
            message = "VPN server status update"
        }

        // Отправляем сообщение администратору
        err := sendMessage(bot, parseChatID(adminChat), message)
        if err != nil {
            c.JSON(500, gin.H{"error": "Failed to send message"})
            return
        }

        c.JSON(200, gin.H{"message": "Update sent"})
    })

    // Маршрут для отправки изображений
    r.GET("/send-photo", func(c *gin.Context) {
        photoURL := c.Query("url")
        if photoURL == "" {
            c.JSON(400, gin.H{"error": "Missing photo URL"})
            return
        }

        // Отправляем изображение администратору
        err := sendPhoto(bot, parseChatID(adminChat), photoURL)
        if err != nil {
            c.JSON(500, gin.H{"error": "Failed to send photo"})
            return
        }

        c.JSON(200, gin.H{"message": "Photo sent"})
    })

    // Запуск веб-сервера
    logger.Fatal(r.Run(":8082"))
}

// parseChatID преобразует строковый ID чата в int64
func parseChatID(chatID string) int64 {
    id, err := strconv.ParseInt(chatID, 10, 64)
    if err != nil {
        logger.Fatalf("Invalid chat ID: %s", chatID)
    }
    return id
}
