package balancer

import (
	"fmt"
)

type Balancer interface {
	Balance() string
}

func New(balancerName string) Balancer {
	fmt.Println(balancerName)
	var newBalancer Balancer
	switch balancerName {
	case "random":
		newBalancer = Random{}
	case "round_robin":
		newBalancer = RoundRobin{}
	}
	return newBalancer
}
