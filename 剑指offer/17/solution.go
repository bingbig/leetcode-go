package solution

func printNumbers(n int) []int {
	if n <= 0 {
		return nil
	}
	max := 1
	for n > 0 {
		max = max * 10
		n--
	}

	tmp := make([]int, max-1)
	for i := 1; i < max; i++ {
		tmp[i-1] = i
	}

	return tmp
}
