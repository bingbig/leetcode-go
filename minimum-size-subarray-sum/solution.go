package solution

func minSubArrayLen(s int, nums []int) int {
	ml := len(nums)
	if ml == 0 {
		return 0
	}

	sum := 0
	p := 0

	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums); i++ {
		if sum >= s {
			for sum >= s && p < i {
				sum -= nums[p]
				p++
			}
			ml = minInt(ml, i-p+1)
		}
		sum += nums[i]
	}

	if ml == len(nums) && sum < s {
		return 0
	}

	if sum >= s {
		for sum >= s && p < len(nums) {
			sum -= nums[p]
			p++
		}

		ml = minInt(ml, len(nums)-p+1)
	}

	return ml
}
