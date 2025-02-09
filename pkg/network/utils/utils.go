package utils

import "fmt"

func LogError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func LogInfo(message string) {
	fmt.Println("INFO:", message)
}
