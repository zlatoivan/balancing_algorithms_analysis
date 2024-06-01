package utils

import (
	"fmt"
	"log"
	"os"
)

func Mean(data []float64) float64 {
	var sum float64
	for _, d := range data {
		sum += d
	}
	return sum / float64(len(data))
}

func GetColorOfBack(b string) int {
	switch b {
	case "1.zlatoivan.ru":
		return 92 // green
	case "2.zlatoivan.ru":
		return 91 // red
	case "3.zlatoivan.ru":
		return 93 // yellow
	}
	return 0
}

func Color(s string, c int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}

func ToLogs(logsCB string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("os.OpenFile: %v", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Printf("file.Close: %v", err)
		}
	}()
	if _, err = file.WriteString(logsCB); err != nil {
		fmt.Printf("file.WriteString: %v", err)
	}
}
