func getLucky(s string, k int) int {
    var numStr string
    for _, ch := range s {
        numStr += strconv.Itoa(int(ch - 'a' + 1))
    }
    var sum int
    for i := 0; i < k; i++ {
        sum = 0
        for _, digit := range numStr {
            sum += int(digit - '0')
        }
        numStr = strconv.Itoa(sum)
    }
    
    return sum
}