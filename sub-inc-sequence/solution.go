package solution

func subIncrSequence(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	start := 0
	length := 0
	p := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			clen := i - p
			if clen > length {
				start = p
				length = clen
			}
			p = i
		}
	}

	clen := len(nums) - p
	if clen > length {
		start = p
		length = clen
	}

	return nums[start : start+length]
}
