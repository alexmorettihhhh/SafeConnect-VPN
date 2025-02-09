// cmd/vpn-server/monitoring.go
package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func getSystemStats() {
	// Получение данных о CPU
	cpus, _ := cpu.Percent(0, false)
	fmt.Printf("CPU Usage: %v\n", cpus)

	// Получение данных о памяти
	vMem, _ := mem.VirtualMemory()
	fmt.Printf("Memory Usage: %v%%\n", vMem.UsedPercent)
}
