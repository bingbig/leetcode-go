package solution

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	maxInt := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	p := 0
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			max = maxInt(max, i-p)
			p = i
		}
	}
	max = maxInt(max, len(nums)-p)

	return max
}
