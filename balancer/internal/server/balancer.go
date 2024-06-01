package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"balancing_algorithms_analysis/internal/server/utils"
)

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

	// В тот кладет новое время
	s.lastTimesBackGr[backend] = append(s.lastTimesBackGr[backend], sec)
	if len(s.lastTimesBackGr[backend]) == 1 {
		m := 0
		for _, v := range s.lastTimesBackGr {
			m = max(m, len(v))
		}
		if m > 1 {
			for i := 0; i < m; i++ {
				s.lastTimesBackGr[backend] = append(s.lastTimesBackGr[backend], sec)
			}
		}
	}
	//fmt.Println("add to main:", sec)
	// Во все остальные копируем последнее время
	for _, k := range s.balancer.Hosts {
		if len(s.lastTimesBackGr[k]) > 0 && k != backend {
			s.lastTimesBackGr[k] = append(s.lastTimesBackGr[k], s.lastTimesBackGr[k][len(s.lastTimesBackGr[k])-1])
			//fmt.Println("add to second:", s.lastTimesBackGr[k][len(s.lastTimesBackGr[k])-1])
		}
	}
	//for k, v := range s.lastTimesBackGr {
	//	fmt.Println(k, v)
	//}

	// Back
	s.lastTimesBack[backend] = append(s.lastTimesBack[backend], sec)
	s.avgTimeBack[backend] = utils.Mean(s.lastTimesBack[backend])

	// All
	s.lastTimesAll = append(s.lastTimesAll, sec)
	s.avgTimeAll = utils.Mean(s.lastTimesAll)

	s.mx.Unlock()
}

func (s *Server) getLog(sec float64, statusCode int, backend string) string {
	secStr := fmt.Sprintf("%.4f", sec)
	status := fmt.Sprintf("%d", statusCode)
	avg := fmt.Sprintf("%.4f", s.avgTimeAll)
	c := utils.GetColorOfBack(backend)
	logs := fmt.Sprintf("balancer choice %s | took %s sec | status %s | average %s sec\n", utils.Color(backend, c), utils.Color(secStr, c), utils.Color(status, c), utils.Color(avg, 96))

	for i, b := range s.balancer.Hosts {
		c = utils.GetColorOfBack(b)
		avg = fmt.Sprintf("%.4f", s.avgTimeBack[b])
		logs += fmt.Sprintf("avg%d %s\n", i+1, utils.Color(avg, c))
	}
	avg = fmt.Sprintf("%.4f", s.avgTimeAll)
	logs += fmt.Sprintf("avgΣ %s\n\n", utils.Color(avg, 96))

	// Черно белые логи для /logs
	logsCB := fmt.Sprintf("balancer choice %s | took %s sec | status %s | average %s sec\n", backend, secStr, status, avg)

	for i, b := range s.balancer.Hosts {
		c = utils.GetColorOfBack(b)
		avg = fmt.Sprintf("%.4f", s.avgTimeBack[b])
		logsCB += fmt.Sprintf("avg%d %s\n", i+1, avg)
	}
	avg = fmt.Sprintf("%.4f", s.avgTimeAll)
	logsCB += fmt.Sprintf("avgΣ %s\n\n", avg)

	utils.ToLogs(logsCB)

	//data := []byte(logsCB)
	//err := os.WriteFile("logs.txt", data, 0644)
	//if err != nil {
	//	fmt.Printf("os.WriteFile: %v", err)
	//}

	//allTms := ""
	//for back, times := range s.lastTimesBack {
	//	tms := ""
	//	for _, tt := range times {
	//		tms += fmt.Sprintf("%.4f ", tt)
	//	}
	//	allTms += fmt.Sprintf("back %s | avg %.4f | times %v\n", back, s.avgTimeBack[back], tms)
	//}

	return logs
}

func (s *Server) ping(w http.ResponseWriter) string {
	backend := s.balancer.ChooseBackend(s.avgTimeBack)

	//s.mx.Lock()
	statusCode, sec := reqAndGetSec(backend)

	s.update(backend, sec)
	//s.mx.Unlock()

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
