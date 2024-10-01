type Trie struct {
	Val   rune
	Count int
	Next  []*Trie
}

func sumPrefixScores(words []string) []int {
	trie := &Trie{}
	for _, word := range words {
		node := trie
		for _, char := range word {
			if node.Next == nil {
				node.Next = make([]*Trie, 26)
			}
			if node.Next[char-'a'] == nil {
				node.Next[char-'a'] = &Trie{Val: char}
			}
			node = node.Next[char-'a']
			node.Count++
		}
	}

	count := make([]int, len(words))
	for i, word := range words {
		node := trie
		for _, char := range word {
			isFound := false
			for _, el := range node.Next {
				if el != nil && el.Val == char {
					count[i] += el.Count
					node = el
					isFound = true
					break
				}
			}
			if !isFound {
				break
			}
		}
	}

	return count
}