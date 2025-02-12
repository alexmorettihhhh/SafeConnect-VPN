package main

import (
    "context"
    "github.com/sirupsen/logrus"
)

// Глобальный логгер
var logger *logrus.Logger

func init() {
    // Инициализация логгера
    logger = logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{}) // Устанавливаем формат JSON для логов
    logger.SetLevel(logrus.InfoLevel)           // Устанавливаем уровень логирования
}

// logError логирует ошибку с дополнительными полями
func logError(ctx context.Context, err error, fields map[string]interface{}) {
    if err != nil {
        fields["error"] = err.Error()
        logger.WithContext(ctx).WithFields(fields).Error("An error occurred")
    }
}

// logInfo логирует информационное сообщение с дополнительными полями
func logInfo(ctx context.Context, message string, fields map[string]interface{}) {
    logger.WithContext(ctx).WithFields(fields).Info(message)
}

func main() {
    // Пример использования
    ctx := context.Background()

    // Логирование информационного сообщения
    logInfo(ctx, "Application started", map[string]interface{}{
        "version": "1.0.0",
        "env":     "production",
    })

    // Логирование ошибки
    err := someFunctionThatMightFail()
    logError(ctx, err, map[string]interface{}{
        "function": "someFunctionThatMightFail",
    })
}

func someFunctionThatMightFail() error {
    // Пример функции, которая может вернуть ошибку
    return nil // или return errors.New("something went wrong")
}
