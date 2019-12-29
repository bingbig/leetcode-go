package main

import (
	"fmt"
)

func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)&1 == 1 {
		return false
	}

	pair := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	ss := []byte(s)
	if _, ok := pair[s[0]]; !ok {
		return false
	}
	l := len(s)
	stack := make([]byte, l/2+1)
	stack[0] = s[0]
	idx := 0
	for i := 1; i < l; i++ {
		_, ok := pair[ss[i]]
		if ok {
			idx++
			stack[idx] = ss[i]
		} else {
			if ss[i] != pair[stack[idx]] {
				return false
			}
			idx--
			if idx < 0 {
				idx = 0
			}
		}
	}

	return idx == 0
}

func main() {
	test := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"[])",
		"[",
	}

	for _, t := range test {
		fmt.Println(t, isValid(t))
	}
}
