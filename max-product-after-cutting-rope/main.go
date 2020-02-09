// 剑指offer 面试题14 p96

package main

import (
	"fmt"
	"math"
)

// 动态规划
func maxProductAfterCuttingRope_solution_1(length int) int {
	if length <= 1 {
		return 0
	}

	if length == 2 {
		return 1
	}

	if length == 3 {
		return 2
	}

	products := make([]int, length+1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	max := 0
	for i := 4; i <= length; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			p := products[j] * products[i-j]
			if max < p {
				max = p
			}
		}
		products[i] = max
	}

	return products[length]
}

// 贪婪算法
// n >= 5时，可以证明
// 2(n-2) > n,  3(n-3)>n，并且 3(n-3) > 2(n -2)
// 因此尽可能多的减长度大于3的绳子可以使得乘积最大
func maxProductAfterCuttingRope_solution_2(length int) int {
	if length <= 1 {
		return 0
	}

	if length == 2 {
		return 1
	}

	if length == 3 {
		return 2
	}

	timesOf3 := length / 3

	// n=4时， 2X2 最大
	if length-timesOf3*3 == 1 {
		timesOf3--
	}

	timesOf2 := (length - timesOf3*3) / 2

	return int(math.Pow(float64(3), float64(timesOf3)) * math.Pow(float64(2), float64(timesOf2)))
}

func main() {
	fmt.Println(maxProductAfterCuttingRope_solution_1(8))
	fmt.Println(maxProductAfterCuttingRope_solution_2(8))
}
