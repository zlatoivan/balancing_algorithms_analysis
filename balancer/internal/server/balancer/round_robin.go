package balancer

import (
	"fmt"
	"math/rand/v2"
)

type RoundRobin struct{}

func (b RoundRobin) Balance() string {
	backend := fmt.Sprintf("%d.zlatoivan.ru", rand.IntN(3))
	fmt.Println("[round robin] balancer have chosen backend: ", backend)
	return backend
}
