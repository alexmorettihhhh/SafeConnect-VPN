package main

/*
#cgo CFLAGS: -g -Wall
#include "traffic.h"
*/
import "C"
import "fmt"

func analyzeTraffic() {
	C.analyze_traffic()
}

func main() {
	analyzeTraffic()
	fmt.Println("Traffic analysis started")
}
