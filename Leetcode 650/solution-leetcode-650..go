func minSteps(n int) int {
    res := 0
    f := 2
    for n > 1 {
        for n % f == 0 {
            res += f
            n /= f
        }
        f++
    }

    return res
}
