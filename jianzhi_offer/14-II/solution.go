package solution

func cuttingRope(n int) int {
	if n < 2 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	prod := make([]int, n+1)
	prod[0] = 0
	prod[1] = 1
	prod[2] = 2
	prod[3] = 3

	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			p := prod[j] * prod[i-j]
			if max < p {
				max = p
			}
		}
		prod[i] = max
	}

	return prod[n]
}
