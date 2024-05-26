package balancer

import (
	"fmt"
	"math/rand/v2"
)

type Random struct{}

func (b Random) Balance() string {
	//backend := "https://localhost:7071"
	backend := fmt.Sprintf("https://%d.zlatoivan.ru", 3+rand.IntN(1))
	return backend
}
