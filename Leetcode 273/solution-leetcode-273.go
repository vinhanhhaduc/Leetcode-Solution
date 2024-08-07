var (
	btwenty = []string{
		"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
		"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
	}
	tens = []string{
		"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
	}
	thousands = []string{
		"", "Thousand", "Million", "Billion",
	}
)

func numberToWords(num int) string {
	if num == 0 {
		return btwenty[0]
	}

	res := ""
	for i := 0; num > 0; i++ {
		if num % 1000 != 0 {
			res = helper(num % 1000) + thousands[i] + " " + res
		}
		num /= 1000
	}
	return strings.TrimSpace(res)
}

func helper(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 {
		return btwenty[num] + " "
	} else if num < 100 {
		return tens[num / 10] + " " + helper(num % 10)
	} else {
		return btwenty[num / 100] + " Hundred " + helper(num % 100)
	}
}