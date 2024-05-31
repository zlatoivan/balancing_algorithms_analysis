package balancer

type RoundRobin struct {
	hosts []string
	last  int
}

func (b RoundRobin) Balance() string {
	backend := b.hosts[b.last]
	b.last = (b.last + 1) % len(b.hosts)
	return backend
}
