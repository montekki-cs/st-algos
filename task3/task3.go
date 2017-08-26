package main

import (
	"./quicksort"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

	tmp_array := make([]int, len(nums), len(nums))
	copy(tmp_array, nums)

	array, cnt := quicksort.QuickSort(tmp_array, quicksort.QUICKSORT_PIVOT_FIRST)
	fmt.Println(cnt)
	fmt.Println(array[:5])

	copy(tmp_array, nums)

	array, cnt = quicksort.QuickSort(tmp_array, quicksort.QUICKSORT_PIVOT_LAST)
	fmt.Println(cnt)
	fmt.Println(array[:5])

	copy(tmp_array, nums)

	array, cnt = quicksort.QuickSort(tmp_array, quicksort.QUICKSORT_PIVOT_MEDIAN)
	fmt.Println(cnt)
	fmt.Println(array[:5])
}
