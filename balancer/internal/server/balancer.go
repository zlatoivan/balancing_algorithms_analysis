package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func mean(data []float64) float64 {
	var sum float64
	for _, d := range data {
		sum += d
	}
	return sum / float64(len(data))
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 96, s)
}

func (s *Server) ping(backend string) {
	start := time.Now()
	client := http.Client{}
	resp, err := client.Get(backend)
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	sec := time.Since(start).Seconds()
	//s.mx.Lock()
	//s.reqTimeArr = append(s.reqTimeArr, sec)
	//s.averageReqTime = mean(s.reqTimeArr)
	//s.mx.Unlock()
	secStr := fmt.Sprintf("%.4f", sec)
	status := fmt.Sprintf("%d", resp.StatusCode)
	avg := fmt.Sprintf("%.4f", s.averageReqTime)
	fmt.Printf("balancer choice %s | took %s sec | status %s | average %s sec\n", green(backend), green(secStr), green(status), blue(avg))
}

func (s *Server) Balancer(_ http.ResponseWriter, _ *http.Request) {
	// здесь клиентом отправить запрос на тот бэкенд, который вернет балансировщик
	backend := s.balancer.Balance()
	go s.ping(backend)
}

func (s *Server) Reload(_ http.ResponseWriter, _ *http.Request) {
	s.averageReqTime = 0
	s.reqTimeArr = []float64{}
	fmt.Println("--- reload ---")
}
