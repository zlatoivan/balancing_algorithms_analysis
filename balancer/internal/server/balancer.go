package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ping(backend string) {
	start := time.Now()
	client := http.Client{}
	resp, err := client.Get(backend)
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	fmt.Printf("Balanser choice: %s. Req took %.4f seconds. Status code = %d\n", backend, time.Since(start).Seconds(), resp.StatusCode)
}

func (s Server) Balancer(w http.ResponseWriter, r *http.Request) {
	// здесь клиентом отправить запрос на тот бэкенд, который вернет балансировщик
	backend := s.balancer.Balance()
	go ping(backend)
}
