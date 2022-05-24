package solution

func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}

	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum+nums[i] < nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		max = MaxInt(sum, max)
	}

	return max
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
