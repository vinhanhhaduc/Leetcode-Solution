func longestValidParentheses(s string) int {
    res := 0
    l, r := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            l++
        } else {
            r++
        }
        
        if l == r {
            if l * 2 > res {
                res = l * 2
            }
        } else if r > l {
            l, r = 0, 0
        }
    }
    
    l, r = 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '(' {
            l++
        } else {
            r++
        }
        
        if l == r {
            if l * 2 > res {
                res = l * 2
            }
        } else if l > r {
            l, r = 0, 0
        }
    }
    
    return res
}