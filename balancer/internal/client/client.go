package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func ping() {
	//start := time.Now()
	client := http.Client{}
	resp, err := client.Get("https://zlatoivan.ru/balancer")
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	defer resp.Body.Close()

	ans := fmt.Sprintf("%d\n", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("io.ReadAll: %v\n", err)
		}
		ans = string(bodyBytes)
	}
	fmt.Printf("%s", ans)

	//fmt.Printf("took %.4f seconds. Status code = %d\n", time.Since(start).Seconds(), resp.StatusCode)
}

func main() {
	start := time.Now()
	n := 10
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			ping()
			wg.Done()
		}()
		//ping()
	}
	wg.Wait()

	//time.Sleep(2 * time.Second)

	sec := time.Since(start).Seconds()
	rps := float64(n) / sec
	fmt.Printf("\nRPS = %.4f\n\n", rps)
}
