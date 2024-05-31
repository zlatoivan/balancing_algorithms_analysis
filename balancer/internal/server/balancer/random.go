package balancer

import (
	"fmt"
	"math/rand/v2"
)

type Random struct{}

func (b Random) Balance() string {
	backend := fmt.Sprintf("%d.zlatoivan.ru", rand.IntN(3)+1)
	return backend
}
