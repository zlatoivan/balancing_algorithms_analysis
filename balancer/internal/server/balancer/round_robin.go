package balancer

import (
	"fmt"
	"math/rand/v2"
)

type RoundRobin struct{}

func (b RoundRobin) Balance() string {
	backend := fmt.Sprintf("%d.zlatoivan.ru", 1+rand.IntN(1))
	fmt.Println("[round robin] balancer have chosen backend: ", backend)
	return backend
}
