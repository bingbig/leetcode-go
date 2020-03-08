package main

func majorityElement(nums []int) int {
	length := len(nums)
	if len(nums) == 0 {
		panic("input must not be empty!")
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maj, count := nums[0], 1
	for i := 1; i < length; i++ {
		if nums[i] == maj {
			count++
		} else {
			count--
			if count <= 0 {
				maj = nums[i]
				count++
			}
		}
	}

	return maj
}
