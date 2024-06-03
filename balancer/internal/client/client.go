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
	//resp, err := client.Get("http://localhost:7070/balancer")
	if err != nil {
		log.Printf("client.Get: %v", err)
	}

	fmt.Printf("took %.4f seconds. Status code = %d\n", time.Since(start).Seconds(), resp.StatusCode)
}

func sequentially(n int) {
	for i := 0; i < n; i++ {
		ping()
	}
}

func parallel(n int) {
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			ping()
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	start := time.Now()
	n := 1

	sequentially(n)

	//parallel(n)

	sec := time.Since(start).Seconds()
	rps := float64(n) / sec
	fmt.Printf("\nRPS = %.4f\n\n", rps)
}
