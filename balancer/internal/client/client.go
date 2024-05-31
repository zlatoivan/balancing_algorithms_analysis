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

	fmt.Printf("%d ", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("io.ReadAll: %v\n", err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
	//fmt.Printf("took %.4f seconds. Status code = %d\n", time.Since(start).Seconds(), resp.StatusCode)
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ping()
			wg.Done()
		}()
	}
	wg.Wait()
	sec := time.Since(start).Seconds()
	rps := 100 / sec
	fmt.Printf("\nRPS = %.4f\n\n", rps)
}
