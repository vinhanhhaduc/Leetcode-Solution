func addOperators(num string, target int) []string {
	var result []string
	if len(num) == 0 {
		return result
	}
	backtrack(num, target, "", 0, 0, 0, &result)
	return result
}
func backtrack(num string, target int, path string, pos int, evaluated int, multed int, result *[]string) {
	if pos == len(num) {
		if evaluated == target {
			*result = append(*result, path)
		}
		return
	}

	for i := pos; i < len(num); i++ {
		if i != pos && num[pos] == '0' {
			break
		}
		currStr := num[pos : i+1]
		curr, _ := strconv.Atoi(currStr)

		if pos == 0 {
			backtrack(num, target, path+currStr, i+1, curr, curr, result)
		} else {
			backtrack(num, target, path + "+" + currStr, i + 1, evaluated + curr, curr, result)
			backtrack(num, target, path + "-" + currStr, i + 1, evaluated - curr, -curr, result)
			backtrack(num, target, path + "*" +currStr, i + 1, evaluated - multed + multed * curr, multed * curr, result)
		}
	}
}