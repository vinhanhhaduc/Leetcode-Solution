func singleNumber(nums []int) []int {
    x := 0
    for _, num := range nums {
        x ^= num
    }
    d := x & (-x)

    num1, num2 := 0, 0
    for _, num := range nums {
        if num & d == 0 {
            num1 ^= num
        } else {
            num2 ^= num
        }
    }
    return []int{num1, num2}
}