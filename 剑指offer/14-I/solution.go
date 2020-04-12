package solution

import "math"

func cuttingRope(n int) int {
	if n <= 1 {
		return 0
	}

	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}
	if n == 4 {
		return 4
	}
	maxRope := make([]int, n+1)
	maxRope[1] = 1
	maxRope[2] = 2
	maxRope[3] = 3
	maxRope[4] = 4

	for length := 5; length <= n; length++ {
		max := 1

		for i := 1; i <= length/2; i++ {
			if maxRope[i]*maxRope[length-i] > max {
				max = maxRope[i] * maxRope[length-i]
			}
		}
		maxRope[length] = max
	}

	return maxRope[n]
}

func cuttingRope2(n int) int {
	if n <= 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	timesOf3 := n / 3

	// n=4时， 2X2 最大
	if n-timesOf3*3 == 1 {
		timesOf3--
	}

	timesOf2 := (n - timesOf3*3) / 2

	return int(math.Pow(float64(3), float64(timesOf3)) * math.Pow(float64(2), float64(timesOf2)))
}
