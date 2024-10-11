func minAddToMakeValid(s string) int {
    open, close := 0, 0

    for _, char := range s {
        if char == '(' {
            open++
        } else {
            if open > 0 {
                open--
            } else {
                close++
            }
        }
    }
    
    return open + close
}