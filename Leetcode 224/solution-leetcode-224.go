func calculate(s string) int {
    stack := []int{}
    result := 0
    currentNumber := 0
    sign := 1
    for i := 0; i < len(s); i++ {
        char := s[i]
        if isDigit(char) {
            currentNumber = currentNumber * 10 + int(char - '0')
        } else if char == '+' {
            result += sign * currentNumber
            currentNumber = 0
            sign = 1
        } else if char == '-' {
            result += sign * currentNumber
            currentNumber = 0
            sign = -1
        } else if char == '(' {
            stack = append(stack, result)
            stack = append(stack, sign)
            result = 0
            sign = 1
        } else if char == ')' {
            result += sign * currentNumber
            result *= stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            result += stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            currentNumber = 0
        }
    }
    result += sign * currentNumber
    return result
}
func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}