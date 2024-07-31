func strStr(haystack string, needle string) int {
    if needle == "" {
		return 0
	}

	j := len(needle)
	i := len(haystack)

	for k := 0; k <= i - j; k++ {
		if haystack[k : k + j] == needle {
			return k
		}
	}
	return -1
}