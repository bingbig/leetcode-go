package solution

func multiply(A int, B int) int {
	if A == 0 || B == 0 {
		return 0
	}
	if A == 1 {
		return B
	}

	if B == 1 {
		return A
	}

	ans := 0
	if A > B {
		ans = multiply(A, B-1) + A
	} else {
		ans = multiply(A-1, B) + B
	}

	return ans
}
