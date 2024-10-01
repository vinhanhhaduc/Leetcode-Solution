func differByOne(word1, word2 string) bool {
	diff := 0
	for i := range word1 {
		if word1[i] != word2[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}
func bfs(beginWord, endWord string, wordSet map[string]bool) map[string][]string {
	queue := []string{beginWord}
	parents := map[string][]string{}
	level := map[string]int{} 
	level[beginWord] = 0
	found := false

	for len(queue) > 0 && !found {
		nextQueue := []string{}
		for _, word := range queue {
			for nextWord := range wordSet {
				if differByOne(word, nextWord) {
					if _, seen := level[nextWord]; !seen {
						level[nextWord] = level[word] + 1
						nextQueue = append(nextQueue, nextWord)
						parents[nextWord] = append(parents[nextWord], word)
						if nextWord == endWord {
							found = true
						}
					} else if level[nextWord] == level[word]+1 {
						parents[nextWord] = append(parents[nextWord], word)
					}
				}
			}
		}
		queue = nextQueue
	}

	return parents
}
func dfs(word string, beginWord string, parents map[string][]string, path []string, results *[][]string) {
	if word == beginWord {
		reversed := make([]string, len(path))
		copy(reversed, path)
		for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
			reversed[i], reversed[j] = reversed[j], reversed[i]
		}
		*results = append(*results, reversed)
		return
	}

	for _, parent := range parents[word] {
		dfs(parent, beginWord, parents, append(path, parent), results)
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if _, exists := wordSet[endWord]; !exists {
		return [][]string{}
	}
	parents := bfs(beginWord, endWord, wordSet)
	results := [][]string{}
	if len(parents) == 0 {
		return results
	}

	path := []string{endWord}
	dfs(endWord, beginWord, parents, path, &results)
	return results
}