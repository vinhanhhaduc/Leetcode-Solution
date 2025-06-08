func isScramble(s1 string, s2 string) bool {
	memo := make(map[string]bool)

	var helper func(i, j, length int) bool
	helper = func(i, j, length int) bool {
		key := fmt.Sprintf("%d-%d-%d", i, j, length)
		if val, exists := memo[key]; exists {
			return val
		}
		if s1[i:i+length] == s2[j:j+length] {
			memo[key] = true
			return true
		}
		if sorted(s1[i:i+length]) != sorted(s2[j:j+length]) {
			memo[key] = false
			return false
		}
		for k := 1; k < length; k++ {
			if helper(i, j, k) && helper(i+k, j+k, length-k) {
				memo[key] = true
				return true
			}
			if helper(i, j+length-k, k) && helper(i+k, j, length-k) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}
	return helper(0, 0, len(s1))
}

func sorted(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}