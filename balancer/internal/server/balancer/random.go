package balancer

import (
	"math/rand/v2"
)

type Random struct {
	Hosts []string
}

func (b *Random) ChooseBackend(_ map[string]float64) string {
	return b.Hosts[rand.IntN(len(b.Hosts))]
}
