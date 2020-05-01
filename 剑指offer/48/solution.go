package solution

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	max := 0
	sub := map[byte]int{}
	for f := 0; f < len(s); f++ {
		if idx, ok := sub[s[f]]; !ok {
			sub[s[f]] = f
		} else {
			if len(sub) > max {
				max = len(sub)
			}
			sub = map[byte]int{}
			f = idx + 1
			sub[s[f]] = f
		}
	}

	if len(sub) > max {
		max = len(sub)
	}

	return max
}
