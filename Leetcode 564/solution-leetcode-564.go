func nearestPalindromic(n string) string {
	num, _ := strconv.ParseInt(n, 10, 64)
	if num <= 10 {
		return strconv.FormatInt(num-1, 10)
	}
	if num == 11 {
		return "9"
	}

	length := len(n)
	left, _ := strconv.ParseInt(n[: (length + 1) / 2], 10, 64)
	a := make([]int64, 5)
	a[0] = helper(left - 1, length % 2 == 0)
	a[1] = helper(left, length % 2 == 0)
	a[2] = helper(left + 1, length % 2 == 0)
	a[3] = int64(math.Pow10(length-1)) - 1
	a[4] = int64(math.Pow10(length)) + 1

	var res int64
	mind := int64(math.MaxInt64)
	for _, i := range a {
		if i == num {
			continue
		}
		d := abs(i - num)
		if d < mind || (d == mind && i < res) {
			mind = d
			res = i
		}
	}

	return strconv.FormatInt(res, 10)
}
func helper(left int64, check bool) int64 {
	palindrome := left
	if !check {
		left /= 10
	}
	for left > 0 {
		palindrome = palindrome*10 + left%10
		left /= 10
	}
	return palindrome
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
