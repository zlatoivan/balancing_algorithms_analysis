package server

import (
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"time"
)

func MultiplyMatrix(n int, matrixA [][]int, matrixB [][]int) [][]int {
	total := 0
	result := genRandMatrix(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				total = total + matrixA[i][k]*matrixB[k][j]
			}
			result[i][j] = total
			total = 0
		}
	}
	return result
}

func genRandMatrix(n int) [][]int {
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = rand.IntN(10)
		}
	}
	return a
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, s)
}

func (s *Server) Duration(w http.ResponseWriter, _ *http.Request) {
	//slp := s.timeSleep
	s.mx.Lock()
	slp := (math.Sin(s.timeSleep) + 1) * 3
	time.Sleep(time.Duration(int64(slp*1000)) * time.Millisecond)
	fmt.Printf("Sleep %s sec\n", green(fmt.Sprintf("%.4f", slp)))
	s.timeSleep += math.Pi / 8
	if s.timeSleep >= 2*math.Pi {
		s.timeSleep -= 2 * math.Pi
	}
	_, err := w.Write([]byte(fmt.Sprintf("%.4f", slp)))
	if err != nil {
		fmt.Printf("w.Write: %v\n", err)
	}
	s.mx.Unlock()
}

//func (s Server) Duration(_ http.ResponseWriter, _ *http.Request) {
//	start := time.Now()
//	n := s.matrixSize
//	matrixA := genRandMatrix(n)
//	matrixB := genRandMatrix(n)
//	MultiplyMatrix(n, matrixA, matrixB)
//	sec := fmt.Sprintf("%.4f", time.Since(start).Seconds())
//	fmt.Printf("Multiply matrices took %s seconds\n", green(sec))
//}
