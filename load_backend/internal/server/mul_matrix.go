package server

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
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

func (s Server) MulMatrices(_ http.ResponseWriter, _ *http.Request) {
	//start := time.Now()
	//n := s.matrixSize
	//matrixA := genRandMatrix(n)
	//matrixB := genRandMatrix(n)
	//MultiplyMatrix(n, matrixA, matrixB)
	//sec := fmt.Sprintf("%.4f", time.Since(start).Seconds())
	//fmt.Printf("Multiply matrices took %s seconds\n", green(sec))

	slp := 1
	time.Sleep(1 * time.Second)
	fmt.Printf("sleep %s sec", green(strconv.Itoa(slp)))
}
