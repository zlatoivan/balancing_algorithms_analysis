package balancer

import (
	"fmt"
	"sync"
)

type RoundRobin struct {
	Hosts []string
	Last  int
	Mx    sync.RWMutex
}

func (b *RoundRobin) ChooseBackend() string {
	b.Mx.Lock()
	backend := b.Hosts[b.Last]
	fmt.Println(b.Last, "  ", (b.Last+1)%len(b.Hosts))
	b.Last = (b.Last + 1) % len(b.Hosts)
	b.Mx.Unlock()

	return backend
}
