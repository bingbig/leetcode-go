package solution

func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ele := start + 2*i
		ans = ans ^ ele
	}
	return ans
}
