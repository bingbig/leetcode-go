package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {

	return false
}

// TODO
func main() {
	test := map[string]string{
		"aa":          "a",
		"aaa":         "a*",
		"ab":          ".*",
		"aab":         "c*a*b",
		"mississippi": "mis*is*p*.",
	}

	for s, p := range test {
		fmt.Println(s, p, isMatch(s, p))
	}
}
