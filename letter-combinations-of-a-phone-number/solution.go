package solution

var nmap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	if len(digits) == 1 {
		res := []string{}
		for _, s := range nmap[digits[0]] {
			res = append(res, string(s))
		}
		return res
	}

	str := nmap[digits[0]]
	coms := letterCombinations(digits[1:])
	res := []string{}
	for _, s := range str {
		for _, c := range coms {
			res = append(res, string(s)+c)
		}
	}

	return res
}
