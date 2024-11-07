func makeFancyString(s string) string {
    result := []byte{}
	for i := 0; i < len(s); i++ {
		if len(result) >= 2 && result[len(result)-1] == s[i] && result[len(result)-2] == s[i] {
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}