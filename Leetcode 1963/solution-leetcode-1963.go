func minSwaps(s string) int {
    balance := 0
    maxImbalance := 0 
    
    for _, char := range s {
        if char == '[' {
            balance++
        } else {
            balance-- 
        }
        if balance < 0 {
            maxImbalance = max(maxImbalance, -balance)
        }
    }
    return (maxImbalance + 1) / 2
}