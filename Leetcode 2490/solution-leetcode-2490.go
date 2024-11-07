func isCircularSentence(sentence string) bool {
    words := strings.Split(sentence, " ")
	if len(words) == 1 {
		return words[0][0] == words[0][len(words[0])-1]
	}
	for i := 0; i < len(words); i++ {
		lastChar := words[i][len(words[i])-1]
		firstChar := words[(i+1)%len(words)][0]
		if lastChar != firstChar {
			return false
		}
	}
	return true
}