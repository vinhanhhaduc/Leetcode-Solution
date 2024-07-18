var phone = []string{'2': "abc",'3': "def",'4': "ghi",'5': "jkl",'6': "mno",'7': "pqrs",'8': "tuv",'9': "wxyz",}
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
		return []string{}
	}
	
	var res []string
	var a []byte
	
	backtrack(0, digits, a, &res)
	return res
}
func backtrack(index int, digits string, a []byte, res *[]string) {
	if index == len(digits) {
		*res = append(*res, string(a))
		return
	}
	
	letters := phone[digits[index]]
	for i := 0; i < len(letters); i++ {
		a = append(a, letters[i])
		backtrack(index + 1, digits, a, res)
		a = a[:len(a) - 1]
	}
}
