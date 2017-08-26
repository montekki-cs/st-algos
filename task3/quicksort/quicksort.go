package quicksort

const (
	QUICKSORT_PIVOT_FIRST = iota
	QUICKSORT_PIVOT_LAST
	QUICKSORT_PIVOT_MEDIAN
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func median(array []int, a, b, c int) int {
	if array[a] > array[b] {
		if array[c] > array[a] {
			return a
		} else {
			if array[c] > array[b] {
				return c
			} else {
				return b
			}
		}
	} else {
		if array[c] > array[b] {
			return b
		} else {
			if array[c] > array[a] {
				return c
			} else {
				return a
			}
		}
	}
}

func quickSortHelper(array []int, count int, pvt int) ([]int, int) {
	var a int
	var cnt = 0

	if len(array) < 2 {
		return array, max(len(array)-1, 0)
	}

	cnt += (len(array) - 1)

	left, right := 0, len(array)-1

	pivot := 0

	switch pvt {
	case QUICKSORT_PIVOT_FIRST:
		pivot = left
	case QUICKSORT_PIVOT_LAST:
		pivot = right
	case QUICKSORT_PIVOT_MEDIAN:
		middle := 0

		if len(array)%2 == 0 {
			middle = len(array)/2 - 1
		} else {
			middle = len(array) / 2
		}

		first, last := 0, len(array)-1

		pivot = median(array, first, middle, last)
	}

	array[0], array[pivot] = array[pivot], array[0]
	i := left + 1

	for j := left + 1; j < len(array); j++ {
		if array[j] < array[0] {
			array[j], array[i] = array[i], array[j]
			i++
		}
	}

	array[0], array[i-1] = array[i-1], array[0]

	_, a = quickSortHelper(array[:i-1], len(array[:i-1]), pvt)
	cnt += a

	_, a = quickSortHelper(array[i:], len(array[i:]), pvt)
	cnt += a

	return array, cnt
}

func QuickSort(array []int, pivotchoose int) ([]int, int) {
	return quickSortHelper(array, 0, pivotchoose)
}
