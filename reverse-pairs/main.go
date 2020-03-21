package main

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 1
		} else {
			return 0
		}
	}

	return mergeSortCount(nums, 0, len(nums)-1)
}

func mergeSortCount(nums []int, start, end int) int {
	if start >= end {
		return 0
	}

	mid := start + (end-start)/2
	count := mergeSortCount(nums, start, mid) + mergeSortCount(nums, mid+1, end)
	tmp := make([]int, end-start+1)
	i, j := start, mid+1
	idx := 0
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp[idx] = nums[i]
			count += j - (mid + 1)
			i++
		} else {
			tmp[idx] = nums[j]
			j++
		}
		idx++
	}

	for ; i <= mid; i++ {
		tmp[idx] = nums[i]
		count += end - (mid + 1) + 1
		idx++
	}

	for ; j <= end; j++ {
		tmp[idx] = nums[j]
		idx++
	}
	copy(nums[start:end+1], tmp)
	return count
}
