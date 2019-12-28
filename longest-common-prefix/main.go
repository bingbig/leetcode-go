package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	num := len(strs)
	if num == 0 {
		return ""
	}

	if num == 1 {
		return strs[0]
	}
	i := 0
	for {
		for j := 1; j < num; j++ {
			if len(strs[j]) < i {
				return strs[j]
			}
			if len(strs[0]) == i || len(strs[j]) == i || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
		i++
	}
}

func main() {
	test := map[string][]string{
		"fl": []string{"flower", "flow", "flight"},
		// "":   []string{"dog", "racecar", "car"},
		"": []string{"", ""},
	}

	for res, strs := range test {
		fmt.Println(strs, res, longestCommonPrefix(strs))
	}
}
