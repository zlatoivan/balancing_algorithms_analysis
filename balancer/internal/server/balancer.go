package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, s)
}

func ping(backend string) {
	start := time.Now()
	client := http.Client{}
	resp, err := client.Get(backend)
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	sec := fmt.Sprintf("%.4f", time.Since(start).Seconds())
	status := fmt.Sprintf("%d", resp.StatusCode)
	fmt.Printf("balancer choice %s | took %s sec | status %s\n", green(backend), green(sec), green(status))
}

func (s Server) Balancer(w http.ResponseWriter, r *http.Request) {
	// здесь клиентом отправить запрос на тот бэкенд, который вернет балансировщик
	backend := s.balancer.Balance()
	go ping(backend)
}
