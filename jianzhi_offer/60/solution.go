package solution

import "math"

func twoSum(n int) []float64 {
	dp := make([][]int, n+1)
	dp[1] = make([]int, 7)
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	for i := 2; i < n+1; i++ {
		dp[i] = make([]int, 6*i+1)
		for j := i; j < 6*i+1; j++ {
			for k := 1; k <= 6; k++ {
				if j-k < i-1 {
					break
				}
				if j-k > 6*(i-1) {
					continue
				}
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}

	total := math.Pow(6, float64(n))
	res := make([]float64, 6*n-n+1)
	for i := n; i <= 6*n; i++ {
		res[i-n] = float64(dp[n][i]) / total
	}

	return res
}
