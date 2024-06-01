package balancer

import (
	"sync"
)

type RoundRobin struct {
	Hosts []string
	Last  int
	mx    sync.Mutex
}

func (b *RoundRobin) ChooseBackend(_ map[string]float64) string {
	b.mx.Lock()
	defer b.mx.Unlock()

	backend := b.Hosts[b.Last]
	//fmt.Println(b.Last)
	b.Last = (b.Last + 1) % len(b.Hosts)

	return backend
}
