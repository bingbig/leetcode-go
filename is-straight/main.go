package main

// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
func isStraight(nums []int) bool {
	pokes := make([]int, 14)
	max, min := 0, 14
	for _, n := range nums {
		if n == 0 {
			continue
		}

		if pokes[n] == 1 {
			return false
		}
		pokes[n] = 1
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max-min < 5
}
