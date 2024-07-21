func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    var res int
    temp := x
    
    for temp > 0 {
        res = res * 10 + temp % 10
        temp /= 10
    }
    
    return res == x
}