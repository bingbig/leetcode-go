package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	n := x
	m := 0
	mirror := 0
	for {
		m = x % 10
		x = x / 10
		mirror = mirror*10 + m
		if x == 0 {
			break
		}
	}

	return mirror == n
}

func main() {
	test := []int{-100, 0, 2, 121, 123454321, 234343}
	for _, num := range test {
		fmt.Println(num, isPalindrome(num))
	}
}
