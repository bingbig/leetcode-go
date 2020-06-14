package solution

func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}

	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) == 1 {
		if p == "." && len(s) == 1 || p == s {
			return true
		} else {
			return false
		}
	}

	if p[1] == '*' {
		if s == "" {
			return isMatch("", p[2:])
		}

		if p[0] == s[0] || p[0] == '.' {
			return isMatch(s[1:], p[2:]) || isMatch(s[1:], p) || isMatch(s, p[2:])
		} else {
			return isMatch(s, p[2:])
		}
	} else if s != "" && (s[0] == p[0] || p[0] == '.') {
		return isMatch(s[1:], p[1:])
	}

	return false
}
