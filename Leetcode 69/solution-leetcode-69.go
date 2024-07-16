func mySqrt(x int) int {
    res := 0
    for res * res < x {
        res++
    }
    if res * res > x {
        res--
    }
    return res
}