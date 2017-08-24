package countinversions

func mergeArrays(array1, array2 []int, count int) ([]int, int) {
	i, j := 0, 0
	size := len(array1) + len(array2)
	res := make([]int, size, size)

	for k := 0; k < len(array1)+len(array2); k++ {
		if i < len(array1) && j < len(array2) {
			if array1[i] < array2[j] {
				res[k] = array1[i]
				i++
				continue
			} else {
				res[k] = array2[j]
				j++
				count += (len(array1) - i)
				continue
			}
		} else if i < len(array1) {
			res[k] = array1[i]
			i++
			continue
		} else if j < len(array2) {
			res[k] = array2[j]
			j++
			continue
		}
	}

	return res, count
}

func CountAndSort(array []int, count int) ([]int, int) {
	if len(array) <= 1 {
		return array, 0
	}

	a_low := array[:len(array)/2]
	a_high := array[len(array)/2:]

	a_low, c1 := CountAndSort(a_low, count)
	a_high, c2 := CountAndSort(a_high, count)

	return mergeArrays(a_low, a_high, c1+c2)
}
