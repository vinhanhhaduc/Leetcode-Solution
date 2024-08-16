func maxDistance(arrays [][]int) int {
    aMin := arrays[0][0]
    aMax := arrays[0][len(arrays[0])-1]
    res := 0
    
    for i := 1; i < len(arrays); i++ {
        res = max(res, abs(aMax - arrays[i][0]))
        res = max(res, abs(arrays[i][len(arrays[i])-1] - aMin))
        aMin = min(aMin, arrays[i][0])
        aMax = max(aMax, arrays[i][len(arrays[i])-1])
    }
        
    return res
}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
