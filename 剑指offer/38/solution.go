package solution

func permutation(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	str := []rune(s)
	var res, ps []string
	for i := 0; i < len(s); i++ {
		str[0], str[i] = str[i], str[0]
		str0 := string(str[0])
		rs := permutation(string(str[1:]))
		for _, r := range rs {
			res = append(res, str0+r)
		}
	}

	tmp := map[string]bool{}
	for _, s := range res {
		if _, ok := tmp[s]; !ok {
			tmp[s] = true
			ps = append(ps, s)
		}
	}

	return ps
}
