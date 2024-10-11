func areSentencesSimilar(sentence1 string, sentence2 string) bool {
     words1 := strings.Fields(sentence1)
    words2 := strings.Fields(sentence2)

    if len(words1) > len(words2) {
        words1, words2 = words2, words1
    }

    i, j := 0, 0
    n1, n2 := len(words1), len(words2)

    for i < n1 && words1[i] == words2[i] {
        i++
    }
    for j < n1-i && words1[n1-1-j] == words2[n2-1-j] {
        j++
    }
    return i+j == n1
}