package main

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	if len(arr) <= k {
		return arr
	}

	numbers := arr[0:k]
	quickSort(numbers, 0, k-1)
	for i := k; i < len(arr); i++ {
		if arr[i] < numbers[k-1] {
			numbers[k-1] = arr[i]
			quickSort(numbers, 0, k-1)
		}
	}

	return numbers
}

func quickSort(data []int, low, high int) {
	if low < high {
		pi := partition(data, low, high)
		quickSort(data, low, pi-1)
		quickSort(data, pi+1, high)
	}
}

func partition(data []int, low, high int) int {
	pivot := data[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if data[j] < pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[high] = data[high], data[i+1]
	return i + 1
}

// heapsort: too slowly
func heapSort(data []int) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		heapify(data, len(data), i)
	}

	for i := len(data) - 1; i >= 0; i-- {
		data[i], data[0] = data[0], data[i]
		heapify(data, i, 0)
	}
}

func heapify(data []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && data[l] > data[largest] {
		largest = l
	}

	if r < n && data[r] > data[largest] {
		largest = r
	}

	if largest != i {
		data[largest], data[i] = data[i], data[largest]
		heapify(data, n, largest)
	}
}
