func countOfAtoms(formula string) string {
    n := len(formula)
	stack := []map[string]int{make(map[string]int)}
	for i := 0; i < n; {
		if formula[i] == '(' {
			stack = append(stack, make(map[string]int))
			i++
		} else if formula[i] == ')' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++
			k := 0
			for i < n && unicode.IsDigit(rune(formula[i])) {
				k = k*10 + int(formula[i]-'0')
				i++
			}
			if k == 0 {
				k = 1
			}
			for elem, count := range top {
				stack[len(stack)-1][elem] += count * k
			}
		} else {
			start := i
			i++
			for i < n && unicode.IsLower(rune(formula[i])) {
				i++
			}
			elem := formula[start:i]
			k := 0
			for i < n && unicode.IsDigit(rune(formula[i])) {
				k = k*10 + int(formula[i]-'0')
				i++
			}
			if k == 0 {
				k = 1
			}
			stack[len(stack)-1][elem] += k
		}
	}

	a := stack[0]
	b := make([]string, 0, len(a))
	for i := range a {
		b = append(b, i)
	}
	sort.Strings(b)

	var res strings.Builder
	for _, i := range b {
		res.WriteString(i)
		count := a[i]
		if count > 1 {
			res.WriteString(fmt.Sprintf("%d", count))
		}
	}
	return res.String()
}