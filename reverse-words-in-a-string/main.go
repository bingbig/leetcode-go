package main

func reverseWords(s string) string {
	var out string
	max := len(s)
	slices := []byte(s)
	end, start := 0, 0
	for end < max {
		for end < max && slices[end] == ' ' {
			end++
		}
		start = end
		if start >= max {
			break
		}

		for end < max && slices[end] != ' ' {
			end++
		}

		if end > start {
			if len(out) != 0 {
				out = " " + out
			}
			out = string(slices[start:end]) + out
		}

	}

	return out
}

// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	if len(s) < n || len(s) <= 1 {
		return s
	}
	sli := []byte(s)
	return string(append(sli[n:], sli[0:n]...))
}
