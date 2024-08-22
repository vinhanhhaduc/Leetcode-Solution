func findComplement(num int) int {
    bitLength := 0
    n := num
    for n > 0 {
        bitLength++
        n >>= 1
    }
    mask := (1 << bitLength) - 1
    return num ^ mask
}
