package solution

func getKthMagicNumber(k int) int {
	if k == 0 {
		return 0
	}

	f3, f5, f7 := 0, 0, 0
	dp := make([]int, k)
	dp[0] = 1
	for i := 1; i < k; i++ {
		dp[i] = minInt(dp[f3]*3, dp[f5]*5, dp[f7]*7)
		if dp[i] == dp[f3]*3 {
			f3++
		}
		if dp[i] == dp[f5]*5 {
			f5++
		}
		if dp[i] == dp[f7]*7 {
			f7++
		}
	}

	return dp[k-1]
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
