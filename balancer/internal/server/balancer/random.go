package balancer

import (
	"math/rand/v2"
)

type Random struct {
	Hosts []string
}

func (b *Random) ChooseBackend() string {
	return b.Hosts[rand.IntN(3)]
}
