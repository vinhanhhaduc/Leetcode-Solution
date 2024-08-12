func findSubstring(s string, words []string) []int {
    if len(words) == 0 || len(s) == 0 {
        return nil
    }
    
    n := len(words[0])
    m := len(words)
    res := []int{}
    wordMap := make(map[string]int)
    for _, word := range words {
        wordMap[word]++
    }
    for i := 0; i < n; i++ {
        left := i
        right := i
        curr := make(map[string]int)
        count := 0
        
        for right + n <= len(s) {
            word := s[right : right  + n]
            right += n
            if _, exists := wordMap[word]; exists {
                curr[word]++
                count++
                for curr[word] > wordMap[word] {
                    leftWord := s[left : left + n]
                    curr[leftWord]--
                    count--
                    left += n
                }
                if count == m {
                    res = append(res, left)
                }
            } else {
                curr = make(map[string]int)
                count = 0
                left = right
            }
        }
    }
    
    return res
}