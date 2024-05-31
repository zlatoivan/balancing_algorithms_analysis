package balancer

import "fmt"

type RoundRobin struct {
	hosts []string
	Last  int
}

func (b RoundRobin) Balance() string {
	backend := b.hosts[b.Last]
	fmt.Println(b.Last, "  ", (b.Last+1)%len(b.hosts))
	b.Last = (b.Last + 1) % len(b.hosts)
	return backend
}
