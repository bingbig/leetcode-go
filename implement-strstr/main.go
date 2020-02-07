package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	if needle == "" {
		return 0
	}

	hl := len(haystack)
	nl := len(needle)
	if nl > hl {
		return -1
	}
	for i := 0; i < hl; i++ {
		for j := 0; j < nl; j++ {
			if i+j >= hl {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break
			}
			if j == nl-1 {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("hello", ""))
	fmt.Println(strStr("hello", "e"))
	fmt.Println(strStr("hello", "w"))
	fmt.Println(strStr("hello", "world"))
	fmt.Println(strStr("mississippi", "issipi"))
}
