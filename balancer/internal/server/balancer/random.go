package balancer

import (
	"math/rand/v2"
)

type Random struct {
	hosts []string
}

func (b Random) Balance() string {
	return b.hosts[rand.IntN(3)]
}
