package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	n := len(nums)
	start := 1
	end := n - 1

	for start < end {
		middle := (end + start) >> 1
		count := 0
		for _, n := range nums {
			if n <= middle {
				count++
			}
		}

		if count > middle {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return start
}

func main() {
	// 1 2 3 1 2 3
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
