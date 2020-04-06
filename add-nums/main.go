package main

// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
func add(a int, b int) int {
	sum := 0
	carry := 0
	for b != 0 {
		sum = a ^ b
		carry = (a & b) << 1
		a = sum
		b = carry
	}

	return a
}
