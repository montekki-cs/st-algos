package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"os"
	"strconv"
	"strings"
	"./unionfind"
)

type Edge struct {
	From	int
	To	int
	Cost	int
}

type EdgeArray []Edge

func (s EdgeArray) Len() int {
	return len(s)
}

func (s EdgeArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s EdgeArray) Less(i, j int) bool {
	return s[i].Cost < s[j].Cost
}

func readGraphFromFile(fname string) (numNodes int, res EdgeArray, err error) {
	var nodesNum int
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return 0, nil, err
	}

	lines := strings.Split(string(b), "\n")

	for i, l := range lines {
		if l == "" {
			continue
		}

		if i == 0 {
			n, err := strconv.Atoi(l)

			if err != nil {
				return 0, nil, err
			}

			nodesNum = n

			res = make([]Edge, 0, 256)
		} else {
			tokens := strings.Split(string(l), " ")

			node1, err := strconv.Atoi(tokens[0])

			if err != nil {
				return 0, nil, err
			}

			node2, err := strconv.Atoi(tokens[1])

			if err != nil {
				return 0, nil, err
			}

			cost, err := strconv.Atoi(tokens[2])

			if err != nil {
				return 0, nil, err
			}

			var e Edge

			e.From, e.To, e.Cost = node1, node2, cost

			res = append(res, e)
		}
	}

	return nodesNum, res, nil
}

func main() {
	nodesNum, graph, err := readGraphFromFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	sort.Sort(graph)

	fmt.Println(nodesNum)
	fmt.Println(graph[:30])

	var u unionfind.Union

	u.Alloc(nodesNum + 1)

	for _, edge := range graph {
		if len(u.ConnComps()) <= 5 {
			break
		}

		if u.Find(edge.From) == u.Find(edge.To) {
			continue
		}

		u.Union(edge.From, edge.To)

		if u.Find(edge.From) != u.Find(edge.To) {
			fmt.Println(edge)
			panic("Asd")
		}
	}

	fmt.Println(u.ConnComps())

	min := 1000000000
	var MinEdge Edge

	for _, e := range graph {
		if u.Find(e.From) != u.Find(e.To) && e.Cost < min {
			MinEdge = e
			min = e.Cost
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println(MinEdge)
	fmt.Println(u.Find(MinEdge.From), " ", u.Find(MinEdge.To))
}
