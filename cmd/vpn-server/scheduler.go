// cmd/vpn-server/scheduler.go
package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func scheduleTasks() {
	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() {
		// Пример задачи: ежедневный перезапуск сервера в 00:00
		fmt.Println("Scheduled task: Restarting VPN server")
		// Здесь ваш код для перезапуска
	})
	c.Start()
}
