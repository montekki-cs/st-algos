package main

import (
	"./countinversions"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func main() {
	nums, err := readFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	_, cnt := countinversions.CountAndSort(nums, 0)
	fmt.Println(cnt)
}
