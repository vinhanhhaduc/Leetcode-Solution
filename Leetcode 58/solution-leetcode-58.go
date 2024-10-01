func lengthOfLastWord(s string) int {
    words := strings.Fields(s)
    res := len(words[len(words)-1])
    return res
}