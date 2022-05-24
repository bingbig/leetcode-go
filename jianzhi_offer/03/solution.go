package solution

func findRepeatNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return n
		} else {
			m[n] = struct{}{}
		}
	}
	return 0
}
