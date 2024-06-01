package balancer

import (
	"sync"
)

type RoundRobin struct {
	Hosts []string
	Last  int
	mx    sync.Mutex
}

func (b *RoundRobin) ChooseBackend() string {
	b.mx.Lock()
	backend := b.Hosts[b.Last]
	//fmt.Println(b.Last)
	b.Last = (b.Last + 1) % len(b.Hosts)
	b.mx.Unlock()

	return backend
}
