func minimumSteps(s string) int64 {
    var res int64 = 0
    var cnt int = 0 
    for i := 0; i < len(s); i++ {
        if s[i] == '0' {
            res += int64(cnt)
        } else {
            cnt++
        }
    }

    return res
}