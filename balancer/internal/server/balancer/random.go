package balancer

import (
	"fmt"
)

type Random struct{}

func (b Random) Balance() string {
	backend := fmt.Sprintf("%d.zlatoivan.ru", 1) // 1 + rand.IntN(2)
	fmt.Println("[random] balancer have chosen backend: ", backend)
	return backend
}
