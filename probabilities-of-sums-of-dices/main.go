package main

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/solution/dong-tai-gui-hua-by-shy-14/
// 0	0
// 1	1 2 3 4 5 6
// 2	  2 3 4 5 6 7 8 9 10 11 12
// 3 	    3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18
// 4	       ...
// n=1，F(1,s)=1,其中s=1,2,3,4,5,6
// n>=2，F(n,s)=F(n-1,s-1)+F(n-1,s-2)+F(n-1,s-3)+F(n-1,s-4)+F(n-1,s-5)+F(n-1,s-6)

func twoSum(n int) []float64 {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 6*n+1)
	}

	for s := 1; s <= 6; s++ {
		dp[1][s] = 1
	}

	for i := 2; i <= n; i++ {
		for s := i; s < 6*n+1; s++ {
			for d := 1; d <= 6; d++ {
				if s-d < i-1 { // first num in row i is i
					break
				}
				dp[i][s] += dp[i-1][s-d]
			}
		}
	}

	total := math.Pow(6, float64(n))
	res := make([]float64, 6*n+1-n)
	for i := n; i < 6*n+1; i++ {
		res[i-n] = float64(dp[n][i]) / total
	}

	return res
}

func main() {
	fmt.Println(twoSum(3))
}
