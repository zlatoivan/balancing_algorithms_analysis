package balancer

import (
	"fmt"
	"math"
	"sort"
	"sync"

	"balancing_algorithms_analysis/internal/server/utils"
)

type WeightedRoundRobin struct {
	Hosts     []string
	Order     []string
	ReqCurNum int
	mx        sync.Mutex
}

func (b *WeightedRoundRobin) ChooseBackend(avgs map[string]float64) string {
	b.mx.Lock()
	//fmt.Println("len(b.Order) =", len(b.Order))
	//fmt.Println("b.ReqCurNum =", b.ReqCurNum)
	if b.ReqCurNum == len(b.Order) {
		weights := make(map[string]int)
		if len(b.Order) == 0 {
			for _, back := range b.Hosts {
				weights[back] = 1
			}
			b.Order = make([]string, len(b.Hosts))
			copy(b.Order, b.Hosts)
			//fmt.Println("ORDER =", b.Order)
		} else {
			m := 0.0
			for _, v := range avgs {
				m = max(m, v)
			}
			// Считаем веса для каждого бэкенда
			for k, v := range avgs {
				weights[k] = int(math.Round(m / v))
				if weights[k] > len(b.Hosts)*2 { // !!! для синуса!
					weights[k] = len(b.Hosts) * 2
				}
			}

			// Сортируем и создаем последовательность вызовов
			keys := make([]string, 0, len(weights))
			for key := range weights {
				keys = append(keys, key)
			}
			comp := func(i, j int) bool {
				return weights[keys[i]] > weights[keys[j]]
			}
			sort.Slice(keys, comp)

			b.Order = make([]string, 0)
			for _, back := range keys {
				weight := weights[back]
				for i := 0; i < weight; i++ {
					b.Order = append(b.Order, back)
				}
			}

			b.ReqCurNum = 0
		}

		logs := "\tNew weights!\n"
		for i, bc := range b.Hosts {
			color := utils.GetColorOfBack(bc)
			weight := fmt.Sprintf("%d", weights[bc])
			logs += fmt.Sprintf("\tw%d %s\n", i+1, utils.Color(weight, color))
		}
		logs += "\n"
		fmt.Printf(logs)

		logsCB := "\tNew weights!\n"
		for i, bc := range b.Hosts {
			logsCB += fmt.Sprintf("\tw%d %d\n", i+1, weights[bc])
		}
		logsCB += "\n"
		utils.ToLogs(logsCB)
	}

	backend := b.Order[b.ReqCurNum]
	b.ReqCurNum++
	b.mx.Unlock()

	return backend
}
