package main

import (
	"./prim"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type job [2]int

type jobs []job
type jobs2 jobs

func (jb jobs) Len() int {
	return len(jb)
}

func (jb jobs) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

func (jb jobs) Less(i, j int) bool {
	ranki := jb[i][0] - jb[i][1]
	rankj := jb[j][0] - jb[j][1]

	if ranki == rankj {
		return jb[i][1]-jb[j][1] <= 0
	} else {
		return ranki < rankj
	}
}

func (jb jobs2) Len() int {
	return len(jb)
}

func (jb jobs2) Swap(i, j int) {
	jb[i], jb[j] = jb[j], jb[i]
}

func (jb jobs2) Less(i, j int) bool {
	ranki := float32(jb[i][0]) / float32(jb[i][1])
	rankj := float32(jb[j][0]) / float32(jb[j][1])

	if ranki == rankj {
		return jb[i][1]-jb[j][1] <= 0
	} else {
		return ranki < rankj
	}
}

func readGraphFromFile(fname string) (res prim.Graph, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return res, err
	}

	lines := strings.Split(string(b), "\n")

	var edgeNum int

	for i, l := range lines {
		if l == "" {
			continue
		}

		tokens := strings.Split(l, " ")

		token1, err := strconv.Atoi(tokens[0])
		if err != nil {
			return res, err
		}

		token2, err := strconv.Atoi(tokens[1])
		if err != nil {
			return res, err
		}

		if i == 0 {
			edgeNum = token2
			res.Edges = make(map[int][]prim.Edge)
			res.Nodes = make(map[int]bool)
		} else {
			token3, err := strconv.Atoi(tokens[2])

			if err != nil {
				return res, err
			}

			edge1 := prim.Edge{token1, token2, token3}
			edge2 := prim.Edge{token2, token1, token3}

			res.Nodes[edge1.From], res.Nodes[edge1.To] = true, true

			_, ok := res.Edges[edge1.From]
			if !ok {
				res.Edges[edge1.From] = make([]prim.Edge, 0, edgeNum)
			}

			res.Edges[edge1.From] = append(res.Edges[edge1.From], edge1)

			_, ok = res.Edges[edge2.From]
			if !ok {
				res.Edges[edge2.From] = make([]prim.Edge, 0, edgeNum)
			}

			res.Edges[edge2.From] = append(res.Edges[edge2.From], edge2)
		}
	}

	return res, err
}

func readJobsFromFile(fname string) (res jobs, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	for i, l := range lines {
		if l == "" {
			continue
		}

		if i == 0 {
			n, err := strconv.Atoi(l)

			if err != nil {
				return nil, err
			}
			res = make([]job, 0, n)
		} else {
			var j job
			tokens := strings.Split(l, " ")

			n1, err := strconv.Atoi(tokens[0])

			if err != nil {
				return nil, err
			}

			n2, err := strconv.Atoi(tokens[1])

			if err != nil {
				return nil, err
			}

			j[0], j[1] = n1, n2

			res = append(res, j)
		}
	}

	return res, err
}

func main() {
	jobs, err := readJobsFromFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	sort.Sort(jobs)

	res := 0
	time := 0
	for _, j := range jobs {
		time += j[0]
		res += time * j[1]
	}

	fmt.Println("res: ", res)

	sort.Sort(jobs2(jobs))

	res = 0
	time = 0
	for _, j := range jobs {
		time += j[0]
		res += time * j[1]
	}

	fmt.Println("res2: ", res)

	graph, err := readGraphFromFile(os.Args[2])

	if err != nil {
		panic(err)
	}

	mst := prim.PrimsMST(graph)
	sum := 0

	for _, v := range mst {
		sum += v.Cost
	}

	fmt.Println("Cost: ", sum)
}
