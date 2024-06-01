package server

import (
	"fmt"
	"html/template"
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

func color(s string, c int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}

func reqAndGetSec(backend string) (int, float64) {
	start := time.Now()
	client := http.Client{}
	resp, err := client.Get("https://" + backend)
	if err != nil {
		log.Printf("client.Get: %v", err)
	}
	sec := time.Since(start).Seconds()

	//defer resp.Body.Close()
	//body := fmt.Sprintf("%d\n", resp.StatusCode)
	//if resp.StatusCode == http.StatusOK {
	//	bodyBytes, err := io.ReadAll(resp.Body)
	//	if err != nil {
	//		log.Printf("io.ReadAll: %v\n", err)
	//	}
	//	body = string(bodyBytes)
	//}
	//sec, _ := strconv.ParseFloat(body, 64)

	return resp.StatusCode, sec
}

func (s *Server) update(backend string, sec float64) {
	s.mx.Lock()
	//for k := range s.lastTimesBack {
	//	if len(s.lastTimesBack[k]) > 0 && k != backend {
	//		s.lastTimesBack[k] = append(s.lastTimesBack[k], s.lastTimesBack[k][len(s.lastTimesBack[k])-1])
	//	}
	//}

	// Back
	s.lastTimesBack[backend] = append(s.lastTimesBack[backend], sec)
	s.avgTimeBack[backend] = mean(s.lastTimesBack[backend])

	// All
	s.lastTimesAll = append(s.lastTimesAll, sec)
	s.avgTimeAll = mean(s.lastTimesAll)
	s.mx.Unlock()
}

func getColorOfBack(b string) int {
	switch b {
	case "1.zlatoivan.ru":
		return 92 // green
	case "2.zlatoivan.ru":
		return 91 // red
	case "3.zlatoivan.ru":
		return 93 // yellow
	}
	return 0
}

func (s *Server) getLog(sec float64, statusCode int, backend string) string {
	secStr := fmt.Sprintf("%.4f", sec)
	status := fmt.Sprintf("%d", statusCode)
	avg := fmt.Sprintf("%.4f", s.avgTimeAll)
	c := getColorOfBack(backend)
	ans := fmt.Sprintf("balancer choice %s | took %s sec | status %s | average %s sec\n", color(backend, c), color(secStr, c), color(status, c), color(avg, 96))

	for _, b := range s.balancer.Hosts {
		c = getColorOfBack(b)
		avg = fmt.Sprintf("%.4f", s.avgTimeBack[b])
		ans += fmt.Sprintf("avg %s\n", color(avg, c))
	}
	avg = fmt.Sprintf("%.4f", s.avgTimeAll)
	ans += fmt.Sprintf("avg %s\n\n", color(avg, 96))

	//allTms := ""
	//for back, times := range s.lastTimesBack {
	//	tms := ""
	//	for _, tt := range times {
	//		tms += fmt.Sprintf("%.4f ", tt)
	//	}
	//	allTms += fmt.Sprintf("back %s | avg %.4f | times %v\n", back, s.avgTimeBack[back], tms)
	//}

	return ans
}

func (s *Server) ping(w http.ResponseWriter) string {
	backend := s.balancer.ChooseBackend()

	statusCode, sec := reqAndGetSec(backend)

	s.update(backend, sec)

	ans := s.getLog(sec, statusCode, backend)

	fmt.Printf(ans)

	//_, err := w.Write([]byte(ans))
	//if err != nil {
	//	fmt.Printf("w.Write: %v\n", err)
	//}

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
