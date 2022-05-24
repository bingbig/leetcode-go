package solution

func countDigitOne(n int) int {
	count, pow := 0, 1
	for n >= pow {
		count += n / (pow * 10) * pow
		if incr := n%(pow*10) - pow + 1; incr > pow {
			count += pow
		} else if incr > 0 {
			count += incr
		}
		pow *= 10
	}
	return count
}
