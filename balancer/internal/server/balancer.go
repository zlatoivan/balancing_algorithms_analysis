package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func green(s string) interface{} {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, s)
}

func ping(backend string) {
	start := time.Now()
	client := http.Client{}
	resp, err := client.Get(backend)
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	fmt.Printf("%s (balancer choice), req took %.4f (sec), %d (status)\n", green(backend), green(time.Since(start).Seconds()), green(resp.StatusCode))
}

func (s Server) Balancer(w http.ResponseWriter, r *http.Request) {
	// здесь клиентом отправить запрос на тот бэкенд, который вернет балансировщик
	backend := s.balancer.Balance()
	go ping(backend)
}
