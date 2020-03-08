package main

import (
	"fmt"
)

func permutation(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}

	mmap := map[string]struct{}{}
	mutations := []string{}
	for i := 0; i < len(s); i++ {
		subs := ""
		if i == 0 {
			subs = s[1:]
		} else if i == len(s)-1 {
			subs = s[0 : len(s)-1]
		} else {
			subs = s[0:i] + s[i+1:]
		}
		ms := permutation(subs)
		for _, m := range ms {
			ns := s[i:i+1] + m
			if _, ok := mmap[ns]; !ok {
				mutations = append(mutations, ns)
				mmap[ns] = struct{}{}
			}
		}
	}

	return mutations
}

// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
func main() {
	s := "aabc"
	fmt.Println(permutation(s))
}
