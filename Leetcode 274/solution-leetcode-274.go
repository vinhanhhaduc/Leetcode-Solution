func hIndex(citations []int) int {
    sort.Ints(citations)
    
    n := len(citations)
    
    for i := 0; i < n; i++ {
        h := n - i
        if citations[i] >= h {
            return h
        }
    }
    
    return 0
}