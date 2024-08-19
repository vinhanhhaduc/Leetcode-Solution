func isPalindrome(s string) bool {
    left, right := 0, len(s) - 1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
func backtrack(s string, start int, path []string, res *[][]string) {
    if start == len(s) {
        temp := make([]string, len(path))
        copy(temp, path)
        *res = append(*res, temp)
        return
    }
    
    for end := start + 1; end <= len(s); end++ {
        substring := s[start:end]
        if isPalindrome(substring) {
            path = append(path, substring)
            backtrack(s, end, path, res)
            path = path[:len(path) - 1]
        }
    }
}
func partition(s string) [][]string {
    var res [][]string
    var path []string
    backtrack(s, 0, path, &res)
    return res
}