package minmaxheap

import "testing"

func max(array []int) int {
	res := array[0]

	for _, val := range array {
		if val > res {
			res = val
		}
	}

	return res
}

func min(array []int) int {
	res := array[0]

	for _, val := range array {
		if val < res {
			res = val
		}
	}

	return res
}

func TestMaxHeap(t *testing.T) {
	cases := []struct {
		array []int
		want  int
	}{
		{[]int{5, 4, 1, 2, 3, 6}, 6},
		{[]int{1, 2, 3, 4, 5, 6}, 6},
		{[]int{7, 8, 1, 4, 2, 3}, 8},
	}
	for _, c := range cases {
		heap := new(Maxheap)
		heap.elements = make([]int, 0, 16)

		for _, e := range c.array {
			heap.Insert(e)
		}

		maxHave := heap.GetMax()

		if maxHave != c.want {
			t.Errorf("MaxHeap got %q, want %q", maxHave, c.want)
		}

		for {
			if len(heap.elements) == 0 {
				break
			}

			maxWant := max(heap.elements)
			maxHave = heap.ExtractMax()

			if maxHave != maxWant {
				t.Errorf("MaxHeap got %d, want %d", maxHave, maxWant)
			}
		}
	}
}

func TestMinHeap(t *testing.T) {
	cases := []struct {
		array []int
		want  int
	}{
		{[]int{5, 4, 1, 2, 3, 6}, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 1},
		{[]int{7, 8, 1, 4, 2, 3}, 1},
	}
	for _, c := range cases {
		heap := new(Minheap)
		heap.elements = make([]int, 0, 16)

		for _, e := range c.array {
			heap.Insert(e)
		}

		minHave := heap.GetMin()

		if minHave != c.want {
			t.Errorf("MaxHeap got %q, want %q", minHave, c.want)
		}

		for {
			if len(heap.elements) == 0 {
				break
			}

			minWant := min(heap.elements)
			minHave = heap.ExtractMin()

			if minHave != minWant {
				t.Errorf("MaxHeap got %d, want %d", minHave, minWant)
			}
		}
	}
}
