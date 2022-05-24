package solution

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}

	i, j, k, max := 1, 0, 0, len(pushed)
	// push first
	stack := make([]int, max)
	stack[0] = pushed[0]

	for i < max {
		if k >= 0 && stack[k] == popped[j] {
			k--
			if k < -1 {
				return false
			}
			j++
		} else {
			k++
			stack[k] = pushed[i]
			i++
		}
	}

	for j < max {
		if k < 0 {
			return false
		}
		if stack[k] == popped[j] {
			k--
			j++
		} else {
			return false
		}
	}

	return k == -1 && j == max
}
