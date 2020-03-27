package main

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
func singleNumber(nums []int) int {
	bitSum := [32]int{}
	for _, n := range nums {
		bitMask := 1
		for j := 0; j < 32; j++ {
			bit := n & bitMask
			if bit != 0 {
				bitSum[j]++
			}
			bitMask = bitMask << 1
		}
	}

	result := 0
	for i := 31; i >= 0; i-- {
		result = result << 1
		result += bitSum[i] % 3
	}

	return result
}
