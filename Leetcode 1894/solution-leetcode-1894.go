func chalkReplacer(chalk []int, k int) int {
    t := 0
    for _, c := range chalk {
        t += c
    }
    k = k % t
    for i, c := range chalk {
        if k < c {
            return i
        }
        k -= c
    }
    return -1
}