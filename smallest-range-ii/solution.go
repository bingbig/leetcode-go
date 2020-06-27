package solution

import (
	"sort"
)

func smallestRangeII(A []int, K int) int {
	sort.Ints(A)

	N := len(A)
	res := A[N-1] - A[0]

	for i := 0; i < N-1; i++ {
		mx := max(A[i]+K, A[N-1]-K)
		mn := min(A[0]+K, A[i+1]-K)
		res = min(res, mx-mn)
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
