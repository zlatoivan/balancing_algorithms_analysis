package balancer

import (
	"fmt"
)

type Random struct{}

func (b Random) Balance() string {
	backend := fmt.Sprintf("https://%d.zlatoivan.ru", 1)
	return backend
}
