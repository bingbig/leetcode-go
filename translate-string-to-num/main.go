package main

// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
func translateNum(num int) int {
	if num < 9 {
		return 1
	}

	tmp := num % 100
	if tmp <= 9 || tmp >= 26 {
		return translateNum(num / 10)
	} else {
		return translateNum(num/100) + translateNum(num/10)
	}
}
