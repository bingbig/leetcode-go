package solution

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxInt := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	max, count := 0, 0
	for _, n := range nums {
		if n == 0 {
			max = maxInt(max, count)
			count = 0
		} else {
			count++
		}
	}
	max = maxInt(max, count)
	return max
}
