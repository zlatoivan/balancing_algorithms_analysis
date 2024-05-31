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

func (s *Server) ping(backend string) *http.Response {
	start := time.Now()
	client := http.Client{}
	resp, err := client.Get("https://" + backend)
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	sec := time.Since(start).Seconds()

	s.mx.Lock()
	s.lastTimesBack[backend] = append(s.lastTimesBack[backend], sec)
	s.avgTimeBack[backend] = mean(s.lastTimesBack[backend])
	//s.lastTimesAll = append(s.lastTimesAll, sec)
	avgs := make([]float64, 0, len(s.avgTimeBack))
	for _, val := range s.avgTimeBack {
		avgs = append(avgs, val)
	}
	s.avgTimeAll = mean(avgs)
	s.mx.Unlock()

	secStr := fmt.Sprintf("%.4f", sec)
	status := fmt.Sprintf("%d", resp.StatusCode)
	avg := fmt.Sprintf("%.4f", s.avgTimeAll)
	fmt.Printf("balancer choice %s | took %s sec | status %s | average %s sec\n", green(backend), green(secStr), green(status), blue(avg))
	return resp
}

func (s *Server) Balancer(w http.ResponseWriter, _ *http.Request) {
	// здесь клиентом отправить запрос на тот бэкенд, который вернет балансировщик
	backend := s.balancer.Balance()
	resp := s.ping(backend)
	w.WriteHeader(resp)
}

func (s *Server) Reload(_ http.ResponseWriter, _ *http.Request) {
	s.avgTimeAll = 0
	//s.lastTimesAll = []float64{}
	fmt.Println("--- reload ---")
}
