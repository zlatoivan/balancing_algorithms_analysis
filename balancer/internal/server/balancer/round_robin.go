package balancer

import "fmt"

type RoundRobin struct {
	hosts []string
	last  int
}

func (b RoundRobin) Balance() string {
	backend := b.hosts[b.last]
	//fmt.Println(b.last, "  ", (b.last+1)%len(b.hosts))
	//b.last = (b.last + 1) % len(b.hosts)
	b.last++
	fmt.Println("last =", b.last)
	return backend
}
