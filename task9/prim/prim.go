package prim

type Edge struct {
	From int
	To   int
	Cost int
}

type Graph struct {
	Nodes map[int]bool
	Edges map[int][]Edge
}

func cheapestEdge(g Graph, x map[int]bool) (res Edge) {
	i := 0
	var cheapest Edge

	for u, _ := range x {
		_, ok := g.Edges[u]

		if ok {
			for _, e := range g.Edges[u] {
				if x[e.From] && !x[e.To] {
					if i == 0 {
						cheapest = e
						i++
					} else {
						if cheapest.Cost > e.Cost {
							cheapest = e
						}
					}
				}
			}
		}
	}

	return cheapest
}

func PrimsMST(g Graph) (T []Edge) {
	X := make(map[int]bool)
	T = make([]Edge, 0)

	for k, _ := range g.Nodes {
		X[k] = true
		break
	}

	for {
		if len(X) == len(g.Nodes) {
			break
		}

		cheapest := cheapestEdge(g, X)

		T = append(T, cheapest)
		X[cheapest.To] = true
	}

	return T
}
