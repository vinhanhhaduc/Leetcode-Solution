func frequencySort(nums []int) []int {
    a := make(map[int]int)
    n := len(nums)
    for i := 0; i < n; i++ {
        a[nums[i]]++
    }
    sort.Slice(nums, func(i, j int)bool{
        if a[nums[i]] == a[nums[j]]{
            return nums[i] > nums[j]
        } else if a[nums[i]] < a[nums[j]] {
            return true
        }
        return false
    })
    return nums
}