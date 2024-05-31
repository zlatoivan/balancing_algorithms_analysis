package balancer

import (
	"fmt"
	"sync"
)

type RoundRobin struct {
	Hosts []string
	Last  int
	mx    sync.RWMutex
}

func (b *RoundRobin) Balance() string {
	b.mx.RLock()
	backend := b.Hosts[b.Last]
	b.mx.RUnlock()

	fmt.Println(b.Last, "  ", (b.Last+1)%len(b.Hosts))

	b.mx.Lock()
	b.Last = (b.Last + 1) % len(b.Hosts)
	b.mx.Unlock()

	return backend
}
