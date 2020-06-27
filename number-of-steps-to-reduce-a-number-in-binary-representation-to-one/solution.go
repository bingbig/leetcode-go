package solution

func numSteps(s string) int {
	carry := 0
	steps := 0
	for len(s) > 1 {
		if carry == 1 {
			if s[len(s)-1] == '1' {
				carry = 1
				s = s[:len(s)-1]
				steps++
			} else {
				carry = 1
				s = s[:len(s)-1]
				steps += 2
			}
		} else {
			if s[len(s)-1] == '0' {
				s = s[:len(s)-1]
				carry = 0
				steps++
			} else {
				carry = 1
				s = s[:len(s)-1]
				steps += 2
			}
		}
	}

	if carry == 1 && s == "1" {
		steps++
	}

	return steps
}
