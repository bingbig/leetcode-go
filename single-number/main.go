package main

import (
	"fmt"
)

// m,m,n
// 0 ^ m = m; m^m = 0; m ^ n ^m = (m ^ m) ^ n = 0 ^ n = n
func singleNumber(nums []int) int {
	a := 0
	for _, n := range nums {
		a ^= n
	}

	return a
}

// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
func main() {
	test := map[int][]int{
		4: []int{4, 1, 2, 1, 2},
		3: []int{2, 2, 1, 1, 4, 4, 3},
	}

	for res, nums := range test {
		fmt.Println(singleNumber(nums), res)
	}
}
