package balancer

type Balancer interface {
	Balance() string
}

func New(balancerName string) Balancer {
	var newBalancer Balancer
	switch balancerName {
	case "random":
		newBalancer = Random{}
	case "round_robin":
		newBalancer = RoundRobin{}
	}
	return newBalancer
}
