package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
func constructArr(a []int) []int {
	if a == nil || len(a) == 0 {
		return a
	}
	n := len(a)
	f, b := make([]int, n), make([]int, n)
	f[0] = 1
	b[n-1] = 1

	for i := 1; i < n; i++ {
		f[i] = f[i-1] * a[i-1]
	}

	for i := n - 2; i >= 0; i-- {
		b[i] = b[i+1] * a[i+1]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = f[i] * b[i]
	}

	return res
}

func main() {
	r := constructArr([]int{7, 2, 2, 4, 2, 1, 8, 8, 9, 6, 8, 9, 6, 3, 2, 1})
	fmt.Println(r)
}
