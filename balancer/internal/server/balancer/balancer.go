package balancer

type Balancer interface {
	Balance() string
}

func New(balancerName string, hosts []string) interface{} {
	var newBalancer interface{}
	switch balancerName {
	case "random":
		newBalancer = Random{hosts: hosts}
	case "round_robin":
		newBalancer = RoundRobin{hosts: hosts, Last: 0}
	case "weighted_round_robin":
		newBalancer = WeightedRoundRobin{hosts: hosts}
	}
	return newBalancer
}
