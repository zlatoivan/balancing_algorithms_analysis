package balancer

import (
	"fmt"
	"math/rand/v2"
)

type WeightedRoundRobin struct {
	Hosts []string
}

func (b *WeightedRoundRobin) ChooseBackend() string {
	backend := fmt.Sprintf("%d.zlatoivan.ru", rand.IntN(1))
	fmt.Println("[round robin] balancer have chosen backend: ", backend)
	return backend
}
