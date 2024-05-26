package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func ping() {
	start := time.Now()
	client := http.Client{}
	resp, err := client.Get("https://zlatoivan.ru/balancer")
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	fmt.Printf("took %.4f seconds. Status code = %d\n", time.Since(start).Seconds(), resp.StatusCode)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			ping()
			wg.Done()
		}()
	}
	wg.Wait()
}
