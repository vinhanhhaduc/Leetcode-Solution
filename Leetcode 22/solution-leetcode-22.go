func backtrack(res *[]string, current string, open, close, max int) {
	if len(current) == 2*max {
		*res = append(*res, current)
		return
	}

	if open < max {
		backtrack(res, current+"(", open+1, close, max)
	}
	if close < open {
		backtrack(res, current+")", open, close+1, max)
	}
}
func generateParenthesis(n int) []string {
	var res []string
	backtrack(&res, "", 0, 0, n)
	return res
}