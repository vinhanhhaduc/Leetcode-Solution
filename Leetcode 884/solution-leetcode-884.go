func uncommonFromSentences(s1 string, s2 string) []string {
    m1, m2 := make(map[string]int), make(map[string]int)

	var uncommon []string

	words1 := strings.Split(s1, " ")
	words2 := strings.Split(s2, " ")

	for _, word := range words1 {
		m1[word]++
	}

	for _, word := range words2 {
		m2[word]++
	}

    for word, count :=range m1 {
        if count == 1 && m2[word]==0 {
            uncommon = append(uncommon, word)
        }
    }

    for word, count :=range m2 {
        if count == 1 && m1[word]==0 {
            uncommon = append(uncommon, word)
        }
    }

	return uncommon
}