package server

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
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

func reqAndGetSec(backend string) (int, float64) {
	//start := time.Now()
	client := http.Client{}
	resp, err := client.Get("https://" + backend)
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	//sec := time.Since(start).Seconds()

	defer resp.Body.Close()
	body := fmt.Sprintf("%d\n", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("io.ReadAll: %v\n", err)
		}
		body = string(bodyBytes)
	}
	fmt.Printf("body = |%s|", body)
	sec, err := strconv.ParseFloat(body, 64)
	if err == nil {
		fmt.Printf("strconv.ParseFloat: %v", err)
	}

	return resp.StatusCode, sec
}

func (s *Server) update(backend string, sec float64) {
	s.mx.Lock()
	s.lastTimesBack[backend] = append(s.lastTimesBack[backend], sec)
	for k := range s.lastTimesBack {
		if k != backend {
			s.lastTimesBack[k] = append(s.lastTimesBack[k], s.lastTimesBack[k][len(s.lastTimesBack[k])-1])
		}
	}
	s.avgTimeBack[backend] = mean(s.lastTimesBack[backend])
	s.lastTimesAll = append(s.lastTimesAll, sec)
	avgs := make([]float64, 0, len(s.avgTimeBack))
	for _, val := range s.avgTimeBack {
		avgs = append(avgs, val)
	}
	s.avgTimeAll = mean(avgs)
	s.mx.Unlock()
}

func getAllTimesStr(lastTimesBack map[string][]float64, avgTimeBack map[string]float64) string {
	allTms := ""
	for back, times := range lastTimesBack {
		tms := ""
		for _, tt := range times {
			tms += fmt.Sprintf("%.4f ", tt)
		}
		allTms += fmt.Sprintf("back %s | avg %.4f | times %v\n", back, avgTimeBack[back], tms)
	}
	return allTms
}

func (s *Server) ping(w http.ResponseWriter) string {
	backend := s.balancer.ChooseBackend()

	statusCode, sec := reqAndGetSec(backend)

	s.update(backend, sec)

	secStr := fmt.Sprintf("%.4f", sec)
	status := fmt.Sprintf("%d", statusCode)
	avg := fmt.Sprintf("%.4f", s.avgTimeAll)
	ans := fmt.Sprintf("balancer choice %s | took %s sec | status %s | average %s sec\n", green(backend), green(secStr), green(status), blue(avg))
	ans += getAllTimesStr(s.lastTimesBack, s.avgTimeBack)
	fmt.Printf(ans)

	_, err := w.Write([]byte(ans))
	if err != nil {
		fmt.Printf("w.Write: %v\n", err)
	}

	return ans
}

func (s *Server) Balancer(w http.ResponseWriter, _ *http.Request) {
	// здесь клиентом отправить запрос на тот бэкенд, который вернет балансировщик

	s.ping(w)

	t, _ := template.ParseFiles("static/template/index.html")
	err := t.Execute(w, "")
	if err != nil {
		log.Printf("t.Execute: %v", err)
	}

	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	s.ping(w)
	//	wg.Done()
	//}()
	//wg.Wait()
}

func (s *Server) Reload(_ http.ResponseWriter, _ *http.Request) {
	s.avgTimeAll = 0
	//s.lastTimesAll = []float64{}
	fmt.Println("--- reload ---")
}
