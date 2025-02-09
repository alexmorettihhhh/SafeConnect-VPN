package main

/*
#cgo CFLAGS: -g -Wall
#include "traffic.h"
*/
import "C"
import "fmt"

// Функция для вызова анализа трафика из C
func analyzeTraffic() {
	C.analyze_traffic()
	fmt.Println("Traffic analysis started.")
}
