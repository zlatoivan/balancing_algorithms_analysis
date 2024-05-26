package server

import (
	"fmt"
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

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("\n%s took %v seconds\n", name, time.Since(start).Seconds())
	}
}

func (s Server) MulMatrices(_ http.ResponseWriter, _ *http.Request) {
	defer timer("MulMatrices")()
	n := s.matrixSize
	matrixA := genRandMatrix(n)
	matrixB := genRandMatrix(n)
	MultiplyMatrix(n, matrixA, matrixB)
}
