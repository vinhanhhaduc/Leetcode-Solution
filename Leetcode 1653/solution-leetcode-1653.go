func minimumDeletions(s string) int {
    count := 0
    b := 0

    for i := 0; i < len(s); i++ {
        if s[i] == 'a' && b > 0 {
            b--
            count++
        } else if s[i] == 'b' {
            b++
        }
    }

    return count
}