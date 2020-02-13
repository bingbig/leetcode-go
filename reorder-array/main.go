package main

import "fmt"

func exchange(nums []int) []int {
	return ReorderArray(nums, isEven)
}

func isEven(n int) bool {
	return (n & 0x1) == 0
}

func ReorderArray(nums []int, fn func(int) bool) []int {
	i := 0
	j := len(nums) - 1

	for i < j {
		for i < j && !fn(nums[i]) {
			i++
		}

		for j > i && fn(nums[j]) {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return nums
}

func main() {
	nums := []int{0, 2}
	fmt.Println(exchange(nums))
}
