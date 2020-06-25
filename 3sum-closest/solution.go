package solution

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	best := math.MaxInt32

	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		j, k := i+1, n-1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s == target {
				return target
			}
			update(s)
			if s < target {
				nj := j + 1
				for nj < k && nums[j] == nums[nj] {
					nj++
				}
				j = nj
			} else {
				nk := k - 1
				for nk > j && nums[k] == nums[nk] {
					nk--
				}
				k = nk
			}
		}
	}
	return best
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
