func longestCommonPrefix(strs []string) string {
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        for len(prefix) > 0 && len(strs[i]) < len(prefix) || len(prefix) > 0 && strs[i][:len(prefix)] != prefix {
            prefix = prefix[:len(prefix)-1]
        }
        if len(prefix) == 0 {
            return ""
        }
    }
    return prefix
}