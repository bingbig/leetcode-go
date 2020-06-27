package solution

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i, _ := range dp {
		dp[i] = max
	}

	dp[0] = 0
	for a := 0; a <= amount; a++ {
		for c := 0; c < len(coins); c++ {
			if coins[c] <= a {
				dp[a] = min(dp[a], dp[a-coins[c]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
