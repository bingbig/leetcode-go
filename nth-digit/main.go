package main

/**
区间	数字个数	字符数个数	字符数累积总和
1-9	9	9*1=9	9
10-99	90	90*2=180	180+9=189
100-999	900	900*3=2700	2700+189=2889
...	9 * 10^(位数-1)	数字个数* 位数	累积和
**/

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
