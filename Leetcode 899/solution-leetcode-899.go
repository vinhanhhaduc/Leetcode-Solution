func orderlyQueue(s string, k int) string {
    if k == 1 {
        ss := s
        for i := 1; i < len(s); i++ {
            n := s[i:] + s[:i]
            if n < ss {
                ss = n
            }
        }
        return ss
    } else {
        chars := []rune(s)
        sort.Slice(chars, func(i, j int) bool {
             return chars[i] < chars[j] 
        })
        return string(chars)
    }
}