package main

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}

	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		max = maxInt(max, sum)
	}
	return max
}
