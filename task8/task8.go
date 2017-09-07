package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

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

func computeTwoSums(nums map[int]bool, t int) int {
	res := 0

	for x, _ := range nums {
		y := t - x

		if y == x {
			continue
		}

		_, ok := nums[y]

		if ok {
			res++
			break
		}
	}

	return res
}

func computeTwoSumsRoutine(nums map[int]bool, tArray []int) int {
	res := 0

	for _, v := range tArray {
		res += computeTwoSums(nums, v)
	}

	return res
}

func main() {
	nums, err := readNumbersFromFile(os.Args[1])

	if err != nil {
		panic(err)
	}

	numsMap := make(map[int]bool, len(nums))

	c := make(chan int, 128)

	for _, v := range nums {
		numsMap[v] = true
	}

	tArray := make([]int, 0, 1000)

	for i := -10000; i <= 10000; i++ {
		if i > -10000 && i%1000 == 0 {
			tArray = append(tArray, i)
			tArrayTmp := make([]int, len(tArray))
			copy(tArrayTmp, tArray)

			wg.Add(1)
			go func(tA []int, ii int, cc chan int) {
				defer wg.Done()
				fmt.Println("Computing in range: ", tA[0], " ", tA[len(tA)-1])
				res := computeTwoSumsRoutine(numsMap, tA)
				cc <- res
			}(tArrayTmp, i, c)
			tArray = make([]int, 0, 1000)
		} else {
			tArray = append(tArray, i)
		}
	}

	wg.Wait()

	res := 0
	close(c)

	for i := range c {
		res += i
	}

	// 427
	fmt.Println("Result: ", res)
}
