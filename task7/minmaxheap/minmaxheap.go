package minmaxheap

type Minheap struct {
	elements []int
}

type Maxheap struct {
	elements []int
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func (heap *Minheap) heapify(idx int) {
	r := right(idx)
	l := left(idx)
	smallest := idx

	if l < len(heap.elements) && heap.elements[l] < heap.elements[idx] {
		smallest = l
	}

	if r < len(heap.elements) && heap.elements[r] < heap.elements[smallest] {
		smallest = r
	}

	if smallest != idx {
		heap.elements[idx], heap.elements[smallest] = heap.elements[smallest], heap.elements[idx]

		heap.heapify(smallest)
	}
}

func (heap *Minheap) Alloc(capacity int) {
	heap.elements = make([]int, 0, capacity)
}

func (heap *Minheap) Size() int {
	return len(heap.elements)
}

func (heap *Minheap) GetMin() int {
	return heap.elements[0]
}

func (heap *Minheap) ExtractMin() int {
	var res int

	if len(heap.elements) == 1 {
		res = heap.elements[0]

		heap.elements = heap.elements[:0]
	} else {
		res = heap.elements[0]
		heap.elements[0] = heap.elements[len(heap.elements)-1]
		heap.elements = heap.elements[:len(heap.elements)-1]

		heap.heapify(0)
	}

	return res
}

func (heap *Minheap) Insert(a int) {
	heap.elements = append(heap.elements, a)

	i := len(heap.elements) - 1

	for {
		if i <= 0 || heap.elements[parent(i)] < heap.elements[i] {
			break
		}

		heap.elements[parent(i)], heap.elements[i] = heap.elements[i], heap.elements[parent(i)]

		i = parent(i)
	}
}

func (heap *Minheap) Delete(a int) {
}

func (heap *Maxheap) heapify(idx int) {
	r := right(idx)
	l := left(idx)
	biggest := idx

	if l < len(heap.elements) && heap.elements[l] > heap.elements[idx] {
		biggest = l
	}

	if r < len(heap.elements) && heap.elements[r] > heap.elements[biggest] {
		biggest = r
	}

	if biggest != idx {
		heap.elements[idx], heap.elements[biggest] = heap.elements[biggest], heap.elements[idx]

		heap.heapify(biggest)
	}
}

func (heap *Maxheap) Alloc(capacity int) {
	heap.elements = make([]int, 0, capacity)
}

func (heap *Maxheap) Size() int {
	return len(heap.elements)
}

func (heap *Maxheap) GetMax() int {
	return heap.elements[0]
}

func (heap *Maxheap) ExtractMax() int {
	var res int

	if len(heap.elements) == 1 {
		res = heap.elements[0]

		heap.elements = heap.elements[:0]
	} else {
		res = heap.elements[0]
		heap.elements[0] = heap.elements[len(heap.elements)-1]
		heap.elements = heap.elements[:len(heap.elements)-1]

		heap.heapify(0)
	}

	return res
}

func (heap *Maxheap) Insert(a int) {
	heap.elements = append(heap.elements, a)

	i := len(heap.elements) - 1

	for {
		if i <= 0 || heap.elements[parent(i)] > heap.elements[i] {
			break
		}

		heap.elements[parent(i)], heap.elements[i] = heap.elements[i], heap.elements[parent(i)]

		i = parent(i)
	}
}

func (heap *Maxheap) Delete(a int) {
}
