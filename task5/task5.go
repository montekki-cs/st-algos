package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type edge struct {
	from, to int
}

type Leader struct {
	leader  int
	populae int
}

type Leaders []Leader

func (slice Leaders) Len() int {
	return len(slice)
}

func (slice Leaders) Less(i, j int) bool {
	return slice[i].populae < slice[j].populae
}

func (slice Leaders) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func readGraphFromFile(fname string) (edges []edge, err error) {
	b, err := ioutil.ReadFile(fname)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	res := make([]edge, 0, len(lines))

	i := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		edgeStr := strings.Split(l, " ")

		n1, err := strconv.Atoi(edgeStr[0])
		if err != nil {
			return nil, err
		}

		n2, err := strconv.Atoi(edgeStr[1])
		if err != nil {
			return nil, err
		}

		var app edge
		app.from, app.to = n1, n2

		res = append(res, app)

		i++
	}

	return res, nil
}

func reverseGraph(edges []edge) (rev []edge) {
	res := make([]edge, len(edges))

	for i, edge := range edges {
		res[i].to, res[i].from = edge.from, edge.to
	}

	return res
}

func edgesToGraph(edges []edge) map[int][]int {
	res := make(map[int][]int)

	for _, edge := range edges {
		_, ok := res[edge.from]

		if !ok {
			res[edge.from] = make([]int, 0, 256)
		}

		res[edge.from] = append(res[edge.from], edge.to)
		_, ok = res[edge.to]

		if !ok {
			res[edge.to] = make([]int, 0, 256)
		}
	}

	return res
}

func dfs(graph map[int][]int, explored map[int]bool, leader map[int]int,
	node int, s int, f map[int]int, t *int) int {
	explored[node] = true
	leader[node] = s

	for _, j := range graph[node] {
		if !explored[j] {
			dfs(graph, explored, leader, j, s, f, t)
		}
	}

	*t++

	f[node] = *t

	return 0
}

func dfsLoop(graph map[int][]int, order []int, stoptimes map[int]int) ([]int, map[int]int) {
	explored := make(map[int]bool)
	f := make(map[int]int)
	leader := make(map[int]int)

	i := 0
	t := 0
	s := 0

	for key := range graph {
		f[key] = 0
		leader[key] = 0
		explored[key] = false
		i++
	}

	for _, i := range order {
		if !explored[i] {
			if stoptimes != nil {
				s = stoptimes[i]
			}
			dfs(graph, explored, leader, i, s, f, &t)
		}
	}

	res := make([]int, len(graph))

	i = 0
	for k, v := range f {
		res[len(res)-v] = k
	}

	if stoptimes != nil {
		leaders := make(map[int]int)

		for _, v := range leader {
			_, ok := leaders[v]
			if !ok {
				leaders[v] = 1
			} else {
				leaders[v]++
			}
		}

		ldrs := make(Leaders, len(leaders))

		i = 0
		for k, v := range leaders {
			ldrs[i].leader, ldrs[i].populae = k, v
			i++
		}

		sort.Sort(ldrs)

		// 211 313 459 968 434821
		result := make([]int, 5)
		i = 0
		for _, val := range ldrs[len(ldrs)-5:] {
			result[i] = val.populae
			i++
		}

		fmt.Println(result)
	}

	return res, f
}

func main() {
	var graph, revGraph []edge
	graph, err := readGraphFromFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	revGraph = reverseGraph(graph)
	fmt.Println(graph[:4])
	fmt.Println(revGraph[:4])

	res := edgesToGraph(graph)
	res2 := edgesToGraph(revGraph)

	order := make([]int, len(res))

	j := 0
	for i := len(res) - 1; i >= 0; i-- {
		order[j] = i + 1
		j++
	}

	order, stoptimes := dfsLoop(res2, order, nil)
	dfsLoop(res, order, stoptimes)
}
