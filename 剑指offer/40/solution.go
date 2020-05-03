package solution

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}

	return quickSort(arr)[:k]
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var l, m, r []int

	pivot := arr[0]
	for _, v := range arr {
		if v < pivot {
			l = append(l, v)
		} else if v == pivot {
			m = append(m, v)
		} else {
			r = append(r, v)
		}
	}

	l = quickSort(l)
	r = quickSort(r)
	return append(append(l, m...), r...)
}
