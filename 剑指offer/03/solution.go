package solution

func findRepeatNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	m := map[int]int{}
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return n
		} else {
			m[n] = 1
		}
	}
	return 0
}
