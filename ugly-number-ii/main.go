package main

func nthUglyNumber(n int) int {
	if n < 1 {
		return 0
	}

	if n <= 5 {
		return n
	}

	uglyArr := make([]int, n)
	uglyArr[0] = 1
	idx := 1
	m2 := 0
	m3 := 0
	m5 := 0

	for idx < n {
		min := minInt(uglyArr[m2]*2, uglyArr[m3]*3, uglyArr[m5]*5)
		uglyArr[idx] = min

		for uglyArr[m2]*2 <= uglyArr[idx] {
			m2++
		}
		for uglyArr[m3]*3 <= uglyArr[idx] {
			m3++
		}
		for uglyArr[m5]*5 <= uglyArr[idx] {
			m5++
		}

		idx++
	}

	return uglyArr[idx-1]
}

func minInt(a, b, c int) int {
	if a > b {
		a, b = b, a
	}

	if a > c {
		a, c = c, a
	}

	return a
}
