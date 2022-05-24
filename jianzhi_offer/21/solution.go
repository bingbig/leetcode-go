package solution

func exchange(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	h, t := 0, len(nums)-1
	for h < t {
		for ; h < t; h++ {
			if isEven(nums[h]) {
				break
			}
		}

		for ; t > h; t-- {
			if !isEven(nums[t]) {
				break
			}
		}
		nums[h], nums[t] = nums[t], nums[h]
		h++
		t--
	}

	return nums
}

func isEven(n int) bool {
	return n&1 == 0
}
