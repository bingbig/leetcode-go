package solution

func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			return i
		} else if nums[i] > i {
			i = nums[i] - 1
		}
	}

	return -1
}
