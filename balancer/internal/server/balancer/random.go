package balancer

import (
	"math/rand/v2"
)

type Random struct {
	hosts []string
}

func (b Random) ChooseBackend() string {
	return b.hosts[rand.IntN(3)]
}
