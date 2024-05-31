package server

import (
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/opts"
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

func genYSin() []opts.LineData {
	y := make([]opts.LineData, 0)
	for i := 0; i < 40; i++ {
		val := math.Sin(float64(i)*math.Pi/8) + 1 // синус
		//val += rand.Float64() / 5                 // шум
		val = val - math.Mod(val, 0.01) // 2 знака после запятой
		y = append(y, opts.LineData{Value: val})
	}
	return y
}

func (s *Server) Duration(_ http.ResponseWriter, _ *http.Request) {
	//slp := s.timeSleep
	slp := s.timeSleep
	time.Sleep(time.Duration(slp) * time.Second)
	fmt.Printf("Sleep %s sec\n", green(fmt.Sprintf("%.4f", slp)))
	s.timeSleep = s.timeSleep + math.Pi/8
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
