func countConsistentStrings(allowed string, words []string) int {
    allowedSet := make(map[rune]bool)
	for _, char := range allowed {
		allowedSet[char] = true
	}

	count := 0
	for _, word := range words {
		check := true
		for _, char := range word {
			if !allowedSet[char] {
				check = false
				break
			}
		}
		if check {
			count++
		}
	}

	return count
}