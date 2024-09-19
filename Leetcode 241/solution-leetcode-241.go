func diffWaysToCompute(expression string) []int {
	memo := make(map[string][]int)
	return helper(expression, memo)
}

func helper(expression string, memo map[string][]int) []int {
	if val, found := memo[expression]; found {
		return val
	}
	
	results := []int{}
	for i := 0; i < len(expression); i++ {
		char := expression[i]
		if char == '+' || char == '-' || char == '*' {
			left := expression[:i]
			right := expression[i+1:]
			leftResults := helper(left, memo)
			rightResults := helper(right, memo)
			for _, l := range leftResults {
				for _, r := range rightResults {
					switch char {
					case '+':
						results = append(results, l+r)
					case '-':
						results = append(results, l-r)
					case '*':
						results = append(results, l*r)
					}
				}
			}
		}
	}
	if len(results) == 0 {
		num, _ := strconv.Atoi(expression)
		results = append(results, num)
	}
	memo[expression] = results
	return results
}