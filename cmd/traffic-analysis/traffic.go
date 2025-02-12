package main

/*
#cgo CFLAGS: -g -Wall
#include "traffic.h"
*/
import "C"
import (
    "errors"
    "flag"
    "fmt"
    "log"
    "os"
)

// Константы для уровней логирования
const (
    LogLevelInfo  = "info"
    LogLevelDebug = "debug"
    LogLevelError = "error"
)

// Глобальные переменные для настройки
var (
    logLevel string
    config   TrafficConfig
)

// TrafficConfig представляет конфигурацию для анализа трафика
type TrafficConfig struct {
    Interface string // Сетевой интерфейс для анализа
    Duration  int    // Продолжительность анализа в секундах
}

func init() {
    // Инициализация флагов командной строки
    flag.StringVar(&logLevel, "log-level", LogLevelInfo, "Уровень логирования (info, debug, error)")
    flag.StringVar(&config.Interface, "interface", "eth0", "Сетевой интерфейс для анализа")
    flag.IntVar(&config.Duration, "duration", 60, "Продолжительность анализа в секундах")
    flag.Parse()

    // Настройка логгера
    setupLogger()
}

func setupLogger() {
    switch logLevel {
    case LogLevelDebug:
        log.SetFlags(log.LstdFlags | log.Lshortfile)
        log.Println("Logging level set to DEBUG")
    case LogLevelError:
        log.SetOutput(os.Stderr)
        log.Println("Logging level set to ERROR")
    default:
        log.SetFlags(log.LstdFlags)
        log.Println("Logging level set to INFO")
    }
}

// analyzeTraffic вызывает C-функцию для анализа трафика
func analyzeTraffic() error {
    log.Printf("Starting traffic analysis on interface %s for %d seconds\n", config.Interface, config.Duration)

    // Проверяем, что сетевой интерфейс указан
    if config.Interface == "" {
        return errors.New("network interface is not specified")
    }

    // Вызываем C-функцию
    result := C.analyze_traffic(C.CString(config.Interface), C.int(config.Duration))
    if result != 0 {
        return fmt.Errorf("C function analyze_traffic failed with code %d", result)
    }

    log.Println("Traffic analysis completed successfully")
    return nil
}

// stopTrafficAnalysis останавливает анализ трафика
func stopTrafficAnalysis() {
    log.Println("Stopping traffic analysis...")
    C.stop_traffic_analysis()
    log.Println("Traffic analysis stopped")
}

func main() {
    // Выводим начальное сообщение
    log.Println("Traffic analysis application started")

    // Запускаем анализ трафика
    err := analyzeTraffic()
    if err != nil {
        log.Fatalf("Failed to analyze traffic: %v", err)
    }

    // Останавливаем анализ трафика при завершении программы
    defer stopTrafficAnalysis()

    log.Println("Application finished successfully")
}
