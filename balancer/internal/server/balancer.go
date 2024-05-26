package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n%s took %v seconds\n", name, time.Since(start).Seconds())
	}
}

func (s Server) Balancer(w http.ResponseWriter, r *http.Request) {
	// здесь клиентом отправить запрос на тот бэкенд, который вернет балансировщик
	backend := s.balancer.Balance()
	client := http.Client{}
	defer timer("Get req to" + backend)
	_, err := client.Get(backend)
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
}
