package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	to     int
	length int
}

type Graph map[int][]Edge

func readGraphFromFile(fname string) (res Graph, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	res = make(Graph)

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		tokens := strings.Split(l, "\t")

		var edgeNum int
		for i := 0; i < len(tokens); i++ {
			if tokens[i] == "" {
				continue
			}

			if i == 0 {
				n, err := strconv.Atoi(tokens[i])

				if err != nil {
					return nil, err
				}

				edgeNum = n
				res[edgeNum] = make([]Edge, 0, 256)
			} else {
				edgeTokens := strings.Split(tokens[i], ",")

				edgeTo, err := strconv.Atoi(edgeTokens[0])

				if err != nil {
					return nil, err
				}

				edgeLen, err := strconv.Atoi(edgeTokens[1])

				if err != nil {
					return nil, err
				}

				res[edgeNum] = append(res[edgeNum], Edge{edgeTo, edgeLen})
			}
		}
	}

	return res, err
}

func dijkstraPath(graph Graph, start int) map[int]int {
	X := make(map[int]bool)
	A := make(map[int]int)

	X[start] = true
	A[start] = 0

	for {
		if len(X) == len(graph) {
			break
		}

		i := 0

		var min Edge
		var vV int = 0

		for v := range X {
			for _, w := range graph[v] {
				_, ok := X[w.to]

				if !ok {
					if i == 0 {
						min = w
						vV = v
					} else {
						if A[v]+w.length < A[vV]+min.length {
							min = w
							vV = v
						}
					}
					i++
				}
			}
		}

		X[min.to] = true
		A[min.to] = A[vV] + min.length
	}

	return A
}

func main() {
	graph, err := readGraphFromFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	lengths := dijkstraPath(graph, 1)

	for _, val := range []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197} {
		fmt.Print(lengths[val], ",")
	}

	fmt.Println("\n")
}
