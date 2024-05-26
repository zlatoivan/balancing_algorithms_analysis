package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		start := time.Now()
		client := http.Client{}
		resp, err := client.Get("https://zlatoivan.ru/balancer")
		if err != nil {
			log.Printf("client.Get: %v", err)
		}
		fmt.Printf("took %.4f seconds. Status code = %d\n", time.Since(start).Seconds(), resp.StatusCode)
	}
}
