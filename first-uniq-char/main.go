package main

func firstUniqChar(s string) byte {
	cc := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		if _, ok := cc[s[i]]; ok {
			cc[s[i]]++
		} else {
			cc[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		if cc[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}
