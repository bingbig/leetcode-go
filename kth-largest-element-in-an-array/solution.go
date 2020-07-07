package solution

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	start, end, target := 0, len(nums)-1, len(nums)-k
	for start <= end {
		p := partition(nums, start, end)
		if p == target {
			return nums[p]
		} else if p < target {
			start = p + 1
		} else {
			end = p - 1
		}
	}

	return -1
}

func partition(nums []int, start, end int) int {
	if end > start {
		rand.Seed(time.Now().UnixNano())
		r := start + rand.Intn(end-start)
		nums[end], nums[r] = nums[r], nums[end]
	}
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}
