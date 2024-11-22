func shortestSubarray(nums []int, k int) int {
    res := math.MaxInt32
    type pair struct {
        prefixSum int64
        endIdx    int
    }
    q := make([]pair, 0)
    
    var curSum int64
    for r := 0; r < len(nums); r++ {
        curSum += int64(nums[r])
        if curSum >= int64(k) {
            res = min(res, r+1)
        }
        for len(q) > 0 && curSum-q[0].prefixSum >= int64(k) {
            res = min(res, r-q[0].endIdx)
            q = q[1:] // pop front
        }
        for len(q) > 0 && q[len(q)-1].prefixSum > curSum {
            q = q[:len(q)-1]
        }
        q = append(q, pair{curSum, r})
    }
    
    if res == math.MaxInt32 {
        return -1
    }
    return res
}