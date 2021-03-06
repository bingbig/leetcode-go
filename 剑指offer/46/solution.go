package solution

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
