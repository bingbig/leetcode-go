package solution

func nthUglyNumber(n int) int {
	if n < 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	time2, time3, time5 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = minInt(dp[time2]*2, dp[time3]*3, dp[time5]*5)
		if dp[i] == dp[time2]*2 {
			time2++
		}
		if dp[i] == dp[time3]*3 {
			time3++
		}
		if dp[i] == dp[time5]*5 {
			time5++
		}
	}

	return dp[n-1]
}

func minInt(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}
