package solution

func findNthDigit(n int) int {
	base, digit := 1, 1
	remove := 9 * base * digit
	for n-remove > 0 {
		n -= remove
		base *= 10
		digit++
		remove = 9 * base * digit
	}

	div, mod := n/digit, n%digit
	loc := base + (div - 1)
	if mod > 0 {
		loc++
		div = digit - mod
		for div > 0 {
			loc = loc / 10
			div--
		}
		return loc % 10
	} else {
		return loc % 10
	}
}
