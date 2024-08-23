func shortestPalindrome(s string) string {
    n := len(s)
    if n == 0 {
        return ""
    }

    revs := reverse(s)
    a := helper(s + "#" + revs)
    add := revs[:n - a[len(a) - 1]]
    
    return add + s
}
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func helper(pattern string) []int {
    a := make([]int, len(pattern))
    length := 0
    i := 1

    for i < len(pattern) {
        if pattern[i] == pattern[length] {
            length++
            a[i] = length
            i++
        } else {
            if length != 0 {
                length = a[length-1]
            } else {
                a[i] = 0
                i++
            }
        }
    }

    return a
}
