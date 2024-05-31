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

func (b *RoundRobin) Balance() string {
	b.Mx.RLock()
	backend := b.Hosts[b.Last]
	b.Mx.RUnlock()

	fmt.Println(b.Last, "  ", (b.Last+1)%len(b.Hosts))

	b.Mx.Lock()
	b.Last = (b.Last + 1) % len(b.Hosts)
	b.Mx.Unlock()

	return backend
}
