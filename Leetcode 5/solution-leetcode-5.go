func longestPalindrome(s string) string {
    res := ""
    st := ""
    for i := 0; i < len(s); i++ {
        for j := i + 1; j < len(s) + 1; j++ {
            st = s[i : j]
            if (isPalindrom(st)) {
                if len(st) > len(res) {
                    res = st
                }
            }
        }
    }
    return res
}
func isPalindrom(str string) bool {
    n := len(str) - 1
	for i := 0; i < n; i++ {
		if str[i] != str[n] {
			return false
		}
		n--
	}
	return true
}