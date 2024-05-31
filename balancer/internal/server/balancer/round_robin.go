package balancer

import "fmt"

type RoundRobin struct {
	Hosts []string
	Last  int
}

func (b *RoundRobin) Balance() string {
	backend := b.Hosts[b.Last]
	fmt.Println(b.Last, "  ", (b.Last+1)%len(b.Hosts))
	b.Last = (b.Last + 1) % len(b.Hosts)
	return backend
}
