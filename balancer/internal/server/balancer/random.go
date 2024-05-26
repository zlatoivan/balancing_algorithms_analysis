package balancer

import (
	"fmt"
	"math/rand/v2"
)

type Random struct{}

func (b Random) Balance() string {
	backend := fmt.Sprintf("https://%d.zlatoivan.ru", 4+rand.IntN(1))
	return backend
}
