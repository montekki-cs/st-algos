package main

import (
	"./minmaxheap"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readNumbersFromFile(fname string) (res []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	res = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		res = append(res, n)
	}

	return res, err
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func computeMedians(nums []int) []int {
	medians := make([]int, 0, len(nums))

	heap := new(minmaxheap.Minheap)
	heap.Alloc(len(nums))

	heap2 := new(minmaxheap.Maxheap)
	heap2.Alloc(len(nums))

	for i, e := range nums {
		if i == 0 {
			heap2.Insert(e)
		} else {
			if e <= heap2.GetMax() {
				heap2.Insert(e)
			} else {
				heap.Insert(e)
			}
		}

		if heap2.Size()-heap.Size() >= 2 && heap2.Size() > 1 {
			tmp := heap2.ExtractMax()
			heap.Insert(tmp)
		} else if heap.Size()-heap2.Size() >= 1 {
			tmp := heap.ExtractMin()
			heap2.Insert(tmp)
		}

		var median int

		median = heap2.GetMax()

		medians = append(medians, median)
	}

	return medians
}

func main() {
	nums, err := readNumbersFromFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	medians := computeMedians(nums)
	res := int64(0)

	for _, e := range medians {
		res += int64(e)
	}

	res %= 10000

	fmt.Println("Result: ", res)
}
