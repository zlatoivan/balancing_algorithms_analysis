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
	backend := b.Hosts[b.Last]
	fmt.Println(b.Last, "  ", (b.Last+1)%len(b.Hosts))
	b.mx.Lock()
	b.Last = (b.Last + 1) % len(b.Hosts)
	b.mx.Lock()
	return backend
}
