package main

import "fmt"

func isNumber(s string) bool {
	switch len(s) {
	case 0:
		return false
	case 1:
		return s[0] >= '0' && s[0] <= '9'
	default:
		if s[0] == ' ' {
			return isNumber(s[1:])
		}
	}

	numeric := false
	s, numeric = scanFloat(s)
	if len(s) == 0 {
		return numeric
	}

	if s[0] == 'e' || s[0] == 'E' {
		if len(s) == 1 || !numeric {
			return false
		}

		s, numeric = scanInt(s[1:])
		if !numeric {
			return false
		}
	}

	for len(s) > 0 {
		if s[0] == ' ' {
			s = s[1:]
		} else {
			return false
		}
	}

	return numeric
}

func scanInt(s string) (string, bool) {
	if len(s) == 0 {
		return "", false
	}

	if len(s) == 1 {
		return "", s[0] >= '0' && s[0] <= '9'
	}

	if s[0] == '+' || s[0] == '-' {
		return scanUint(s[1:])
	}

	return scanUint(s)
}

func scanUfloat(s string) (string, bool) {
	s, numeric := scanUint(s)
	if len(s) == 0 {
		return "", numeric
	}

	if s[0] == '.' {
		if len(s) == 1 {
			return "", numeric
		}
		numeric2 := false
		s, numeric2 = scanUint(s[1:])
		numeric = numeric || numeric2
		if !numeric2 && !numeric {
			return s, false
		}
		if len(s) == 0 {
			return "", true
		}
		numeric = numeric || numeric2
	}

	return s, numeric
}

func scanUint(s string) (string, bool) {
	if len(s) == 0 {
		return "", false
	}

	if len(s) == 1 {
		if s[0] >= '0' && s[0] <= '9' {
			return "", true
		} else {
			return s, false
		}
	}

	count := 0
	for count < len(s) {
		if s[count] >= '0' && s[count] <= '9' {
			count++
		} else {
			break
		}
	}

	if count == len(s) {
		return "", true
	} else {
		return s[count:], count > 0
	}
}

func scanFloat(s string) (string, bool) {
	if len(s) == 0 {
		return "", false
	}

	if len(s) == 1 {
		return "", s[0] >= '0' && s[0] <= '9'
	}

	if s[0] == '+' || s[0] == '-' {
		return scanUfloat(s[1:])
	}

	return scanUfloat(s)
}

func main() {
	fmt.Println(1, isNumber("0") == true)
	fmt.Println(2, isNumber(" 0.1 ") == true)
	fmt.Println(3, isNumber("1 a") == false)
	fmt.Println(4, isNumber("2e10") == true)
	fmt.Println(5, isNumber(" -90e3   ") == true)
	fmt.Println(6, isNumber(" -90e 3   ") == false)
	fmt.Println(7, isNumber(" -90 e3   ") == false)
	fmt.Println(8, isNumber(" 1e") == false)
	fmt.Println(9, isNumber(" e3") == false)
	fmt.Println(10, isNumber("  6e-1") == true)
	fmt.Println(11, isNumber("  99e2.5 ") == false)
	fmt.Println(12, isNumber("53.5e93") == true)
	fmt.Println(13, isNumber(" --6") == false)
	fmt.Println(14, isNumber("-+3") == false)
	fmt.Println(15, isNumber("95a54e53") == false)
	fmt.Println(16, isNumber("0..") == false)
	fmt.Println(17, isNumber(".1") == true)
	fmt.Println(18, isNumber("..1") == false)
	fmt.Println(19, isNumber(".9 ") == true)
	fmt.Println(20, isNumber("6e6.5") == false)
}
