package balancer

import (
	"fmt"
	"math/rand/v2"
)

type Random struct{}

func (b Random) Balance() string {
	backend := fmt.Sprintf("https://%d.zlatoivan.ru", rand.IntN(3))
	return backend
}
