package solution

func merge(A []int, m int, B []int, n int) {
	if n == 0 {
		return
	}

	i := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if A[m] >= B[n] {
			A[i] = A[m]
			m--
		} else {
			A[i] = B[n]
			n--
		}
		i--
	}

	for m >= 0 {
		A[i] = A[m]
		m--
		i--
	}
	for n >= 0 {
		A[i] = B[n]
		n--
		i--
	}
}
