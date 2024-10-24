func maxUniqueSplit(s string) int {
    used := make(map[string]bool)
    return backtrack(s, 0, used)
}
func backtrack(s string, start int, used map[string]bool) int {
    if start == len(s) {
        return 0
    }
    
    res := 0
    
    for i := start; i < len(s); i++ {
        substr := s[start:i+1]
        
        if !used[substr] { 
            used[substr] = true   
            res = max(res, 1 + backtrack(s, i+1, used))
            used[substr] = false 
        }
    }
    
    return res
}