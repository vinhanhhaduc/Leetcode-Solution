func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	result := make([]int, m+n)

	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	num1, num2 = reverse(num1), reverse(num2)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			digit1 := int(num1[i] - '0')
			digit2 := int(num2[j] - '0')
			result[i+j] += digit1 * digit2
			result[i+j+1] += result[i+j] / 10
			result[i+j] %= 10
		}
	}
	i := m + n - 1
	for i > 0 && result[i] == 0 {
		i--
	}
	res := ""
	for ; i >= 0; i-- {
		res += string(result[i] + '0')
	}

	return res
}