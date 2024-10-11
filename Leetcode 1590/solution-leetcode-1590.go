func minSubarray(nums []int, p int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }

    target := totalSum % p
    if target == 0 {
        return 0
    }

    prefixSum := 0
    minLength := len(nums)
    prefixMap := map[int]int{0: -1} 

    for i, num := range nums {
        prefixSum = (prefixSum + num) % p
        need := (prefixSum - target + p) % p
        if prevIndex, found := prefixMap[need]; found {
            minLength = min(minLength, i - prevIndex)
        }
        prefixMap[prefixSum] = i
    }

    if minLength == len(nums) {
        return -1
    }
    return minLength
}