package main

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
func singleNumbers(nums []int) []int {
	s := 0
	for _, n := range nums {
		s ^= n
	}

	k := s & (-s)
	rtn := make([]int, 2)
	for _, n := range nums {
		if n&k == 0 {
			rtn[0] ^= n
		} else {
			rtn[1] ^= n
		}
	}

	return rtn
}
