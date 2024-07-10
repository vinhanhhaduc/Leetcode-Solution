func calculate(s string) int {
	var prev, cur, res int
	c := '+'

	for i, j := range s {
		if unicode.IsDigit(j) {
			cur = cur * 10 + int(j - '0')
		}
		if j == '+' || j == '-' || j == '*' || j == '/' || i == len(s) - 1 {
			switch c {
			case '+':
				res += prev
				prev = cur
			case '-':
				res += prev
				prev = -cur
			case '*':
				prev *= cur
			case '/':
				prev /= cur
			}
			c = j
			cur = 0
		}
	}

	res += prev
	return res
}