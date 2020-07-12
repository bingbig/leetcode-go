package solution

func findMinFibonacciNumbers(k int) int {
	fs := []int{1, 1}
	f1, f2 := 1, 1
	for f2 <= k {
		fs = append(fs, f1+f2)
		f1, f2 = f2, f1+f2
	}

	ans := 0
	i := len(fs) - 2
	for i >= 0 {
		k = k - fs[i]
		ans++
		if k == 0 {
			return ans
		} else if k > 0 {
			for k < fs[i] {
				i--
			}
		} else {
			return -1
		}
	}

	return -1
}
