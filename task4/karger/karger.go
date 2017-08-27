package karger

import (
	"math/rand"
)

func Karger(graph map[int][]int) map[int][]int {
	keys := make([]int, len(graph))

	i := 0
	for k := range graph {
		keys[i] = k
		i++
	}

	for len(graph) > 2 {
		vIdxU := rand.Intn(len(keys))

		idx := keys[vIdxU]

		vIdxV := rand.Intn(len(graph[idx]))

		idx2 := graph[idx][vIdxV]

		for idx, k := range keys {
			if k == idx2 {
				vIdxV = idx
				break
			}
		}

		seen := make(map[int]bool)

		for _, k := range graph[idx] {
			seen[k] = true
		}

		for _, k := range graph[idx2] {
			graph[idx] = append(graph[idx], k)
		}

		for k := range graph {
			var res []int

			for _, m := range graph[k] {
				if m == idx2 {
					if k != idx {
						res = append(res, idx)
					}
				} else {
					if m != k {
						res = append(res, m)
					}
				}

			}

			graph[k] = res
		}

		delete(graph, keys[vIdxV])

		if vIdxV < len(keys) {
			keys = append(keys[:vIdxV], keys[vIdxV+1:]...)
		} else {
			keys = keys[:vIdxV]
		}

	}

	return graph
}
