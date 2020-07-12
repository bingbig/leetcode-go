package solution

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	if len(num2) == 0 {
		return num1
	}

	carry := 0
	sum := make([]byte, len(num1)+1)

	j := len(num1)
	for i := len(num2) - 1; i >= 0; i-- {
		sum[j], carry = addByte(num1[j-1], num2[i], carry)
		j--
	}

	for j > 0 {
		sum[j], carry = addByte(num1[j-1], '0', carry)

		j--
	}

	if carry == 1 {
		sum[j] = '1'
	} else {
		sum = sum[1:]
	}

	return string(sum)
}

func addByte(a, b byte, c int) (byte, int) {
	sum := int(a-'0'+b-'0') + c

	if sum >= 10 {
		return byte(sum-10) + '0', 1
	}
	return byte(sum) + '0', 0
}
