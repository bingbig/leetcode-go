package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}

	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) == 1 {
		if p[0] == '.' && len(s) == 1 || p == s {
			return true
		} else {
			return false
		}
	}

	if p[1] == '*' {
		if s == "" {
			return isMatch("", p[2:])
		}
		if p[0] == s[0] || (p[0] == '.' && s != "") {
			// move to next state
			return isMatch(s[1:], p[2:]) ||
				// stay current state
				isMatch(s[1:], p) ||
				// ignore a "*"
				isMatch(s, p[2:])

		} else {
			return isMatch(s, p[2:])
		}
	} else if (s != "" && p[0] == s[0]) || (p[0] == '.' && s != "") {
		return isMatch(s[1:], p[1:])
	}

	return false
}

func main() {
	fmt.Println(isMatch("a", ".*..a*"))
	fmt.Println(isMatch("ab", ".*c"))
	fmt.Println(isMatch("a", "ab*"))
	fmt.Println(isMatch("caab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("mississippi", "mis*is*ip*."))
}
