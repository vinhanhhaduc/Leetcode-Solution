func ladderLength(beginWord string, endWord string, wordList []string) int {
    wordSet := make(map[string]bool)
    for _, word := range wordList {
        wordSet[word] = true
    }
    if !wordSet[endWord] {
        return 0
    }
    queue := [][]string{{beginWord}}
    visited := make(map[string]bool)
    visited[beginWord] = true
    level := 1

    for len(queue) > 0 {
        nextQueue := [][]string{}
        for _, path := range queue {
            word := path[len(path)-1]
            for i := 0; i < len(word); i++ {
                for c := 'a'; c <= 'z'; c++ {
                    if byte(c) == word[i] {
                        continue
                    }
                    nextWord := word[:i] + string(c) + word[i+1:]
                    if nextWord == endWord {
                        return level + 1
                    }
                    if wordSet[nextWord] && !visited[nextWord] {
                        visited[nextWord] = true
                        newPath := append([]string{}, path...)
                        newPath = append(newPath, nextWord)
                        nextQueue = append(nextQueue, newPath)
                    }
                }
            }
        }
        queue = nextQueue
        level++
    }

    return 0
}