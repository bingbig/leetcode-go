package solution

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	half := myPow(x, n>>1)
	result := half * half
	if n&1 == 1 {
		result *= x
	}

	return result
}
