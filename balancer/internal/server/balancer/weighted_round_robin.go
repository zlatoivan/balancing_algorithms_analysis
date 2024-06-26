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
	Weights   map[string]int
}

func (b *WeightedRoundRobin) ChooseBackend(avgs map[string]float64) string {
	b.mx.Lock()
	defer b.mx.Unlock()

	//fmt.Println("len(b.Order) =", len(b.Order))
	//fmt.Println("b.ReqCurNum =", b.ReqCurNum)
	if b.ReqCurNum == len(b.Order) {
		b.Weights = make(map[string]int)
		if len(b.Order) == 0 {
			for _, back := range b.Hosts {
				b.Weights[back] = 1
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
				b.Weights[k] = int(math.Round(m / v))
				if b.Weights[k] > len(b.Hosts)*2 { // !!! для синуса!
					b.Weights[k] = len(b.Hosts) * 2
				}
			}

			// Сортируем и создаем последовательность вызовов
			keys := make([]string, 0, len(b.Weights))
			for key := range b.Weights {
				keys = append(keys, key)
			}
			comp := func(i, j int) bool {
				return b.Weights[keys[i]] > b.Weights[keys[j]]
			}
			sort.Slice(keys, comp)

			b.Order = make([]string, 0)
			for _, back := range keys {
				weight := b.Weights[back]
				for i := 0; i < weight; i++ {
					b.Order = append(b.Order, back)
				}
			}

			b.ReqCurNum = 0
		}

		logs := ""
		for i, bc := range b.Hosts {
			color := utils.GetColorOfBack(bc)
			weight := fmt.Sprintf("%d", b.Weights[bc])
			logs += fmt.Sprintf("\tupd w%d %s\n", i+1, utils.Color(weight, color))
		}
		fmt.Printf(logs)

		logsCB := ""
		for i, bc := range b.Hosts {
			logsCB += fmt.Sprintf("\tupd w%d %d\n", i+1, b.Weights[bc])
		}
		utils.ToLogs(logsCB)
	}

	backend := b.Order[b.ReqCurNum]
	b.ReqCurNum++

	return backend
}
