func fractionAddition(expression string) string {
	num, d := 0, 1
	i := 0
	n := len(expression)

	for i < n {
		sign := 1
		if expression[i] == '-' {
			sign = -1
			i++
		} else if expression[i] == '+' {
			i++
		}
		numStart := i
		for i < n && expression[i] != '/' {
			i++
		}
		currNum, _ := strconv.Atoi(expression[numStart:i])
		i++
		denomStart := i
		for i < n && expression[i] >= '0' && expression[i] <= '9' {
			i++
		}
		curr, _ := strconv.Atoi(expression[denomStart:i])

		currNum *= sign

		num = num * curr + currNum * d
		d *= curr

		g := gcd(abs(num), abs(d))
		num /= g
		d /= g
	}
	return fmt.Sprintf("%d/%d", num, d)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func gcd(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}