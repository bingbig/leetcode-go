package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := findFirstNum(nums, target, 0, len(nums)-1)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findLastNum(nums, target, 0, len(nums)-1)

	return []int{first, last}
}

func findFirstNum(nums []int, target, start, end int) int {
	if end < start {
		return -1
	}

	mid := start + (end-start)/2
	if nums[mid] < target {
		return findFirstNum(nums, target, mid+1, end)
	} else if nums[mid] == target {
		if mid == 0 {
			return 0
		}
		if nums[mid-1] == target {
			return findFirstNum(nums, target, start, mid-1)
		} else {
			return mid
		}
	} else {
		return findFirstNum(nums, target, start, mid-1)
	}
}

func findLastNum(nums []int, target, start, end int) int {
	if end < start {
		return -1
	}

	mid := start + (end-start)/2
	if nums[mid] < target {
		return findLastNum(nums, target, mid+1, end)
	} else if nums[mid] == target {
		if mid == len(nums)-1 {
			return mid
		}
		if nums[mid+1] == target {
			return findLastNum(nums, target, mid+1, end)
		} else {
			return mid
		}
	} else {
		return findLastNum(nums, target, start, mid-1)
	}
}
