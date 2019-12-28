package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	idx := len(digits) - 1
	for {
		if idx == 0 {
			if digits[0] == 9 {
				rtn := make([]int, len(digits)+1)
				rtn[0] = 1
				return rtn
			}
			digits[0] += 1
			return digits
		}

		if digits[idx] != 9 {
			digits[idx] += 1
			return digits
		}
		digits[idx] = 0
		idx--
	}
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{8, 9, 9}))
	fmt.Println(plusOne([]int{8, 9, 9, 9}))
}
