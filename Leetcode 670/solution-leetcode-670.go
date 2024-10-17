func maximumSwap(num int) int {
    digits := []byte(strconv.Itoa(num))
    last := make([]int, 10)
    for i := 0; i < len(digits); i++ {
        last[digits[i]-'0'] = i
    }
    for i := 0; i < len(digits); i++ {
        for d := 9; d > int(digits[i]-'0'); d-- {
            if last[d] > i {
                digits[i], digits[last[d]] = digits[last[d]], digits[i]
                result, _ := strconv.Atoi(string(digits))
                return result
            }
        }
    }
    return num
}