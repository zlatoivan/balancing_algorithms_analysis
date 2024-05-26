package balancer

import (
	"fmt"
	"math/rand/v2"
)

type Random struct{}

func (b Random) Balance() string {
	backend := fmt.Sprintf("%d.zlatoivan.ru", 1+rand.IntN(0))
	fmt.Println("[random] balancer have chosen backend: ", backend)
	return backend
}
