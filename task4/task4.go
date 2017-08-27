package main

import (
	"./karger"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readGraphFromFile(fname string) (res map[int][]int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	res = make(map[int][]int)

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		vertex := strings.Split(l, "\t")

		vertnum := 0

		for i := 0; i < len(vertex); i++ {
			if vertex[i] == "" {
				continue
			}

			n, err := strconv.Atoi(vertex[i])
			if err != nil {
				return nil, err
			}

			if i == 0 {
				res[n] = make([]int, 0, 256)
				vertnum = n
			} else {
				res[vertnum] = append(res[vertnum], n)
			}
		}
	}

	return res, nil
}

func copyGraph(graph map[int][]int) map[int][]int {
	ret := make(map[int][]int)

	for k := range graph {
		ret[k] = make([]int, len(graph[k]))
		for idx, val := range graph[k] {
			ret[k][idx] = val
		}
	}

	return ret
}

func main() {
	nums, err := readGraphFromFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	var resGraph map[int][]int

	rand.Seed(time.Now().UTC().UnixNano())

	oldLen, newLen := 1000000, 0

	for i := 0; i < len(nums)*2; i++ {
		copyNums := copyGraph(nums)
		newGraph := karger.Karger(copyNums)

		for k := range newGraph {
			newLen = len(newGraph[k])
			break
		}

		if newLen < oldLen {
			oldLen = newLen
			resGraph = copyGraph(newGraph)
		}
	}

	fmt.Println(resGraph)

	for k := range resGraph {
		fmt.Println(len(resGraph[k]))
	}
}
