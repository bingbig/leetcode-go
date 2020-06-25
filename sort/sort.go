package sort

func QuickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)>>1]
		for i < j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}

			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		QuickSort(arr, start, j)
		QuickSort(arr, i, end)
	}
}

func QuickSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, mid, right := make([]int, 0), make([]int, 0), make([]int, 0)
	p := arr[0]
	for _, a := range arr {
		if a < p {
			left = append(left, a)
		} else if a == p {
			mid = append(mid, a)
		} else {
			right = append(right, a)
		}
	}

	left = QuickSort2(left)
	right = QuickSort2(right)

	return append(append(left, mid...), right...)
}

func QuickSort3(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[start]
		for i < j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}

			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		arr[i] = key

		QuickSort(arr, start, i-1)
		QuickSort(arr, i+1, end)
	}
}

func BubbleSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	max := len(arr) - 1
	for i := max; i > 1; i-- {
		for j := 1; j <= i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func InsertSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		max := i + 1
		for j := i; j >= 0; j-- {
			if arr[j] > arr[max] {
				arr[j], arr[max] = arr[max], arr[j]
				max = j
			} else {
				break
			}
		}
	}
}

func SelectionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

}
