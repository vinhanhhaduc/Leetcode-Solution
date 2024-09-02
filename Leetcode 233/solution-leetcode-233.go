func countDigitOne(n int) int {
    count := 0
    factor := 1
    
    for n / factor > 0 {
        l := n - (n / factor) * factor
        c := (n / factor) % 10
        h := n / (factor * 10)
        if c == 0 {
            count += h * factor
        } else if c == 1 {
            count += h * factor + l + 1
        } else {
            count += (h + 1) * factor
        }
        
        factor *= 10
    }
    
    return count
}