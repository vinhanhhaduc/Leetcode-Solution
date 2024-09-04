func isValid(code string) bool {
    stack := []string{}
    i := 0
    n := len(code)
    
    for i < n {
        if i > 0 && len(stack) == 0 {
            return false
        }

        if i+9 <= n && code[i:i+9] == "<![CDATA[" {
            j := i + 9
            for j < n && j+3 <= n && code[j:j+3] != "]]>" {
                j++
            }
            if j+3 > n {
                return false
            }
            i = j + 3
        } else if i+2 <= n && code[i:i+2] == "</" {
            j := i + 2
            for j < n && code[j] != '>' {
                j++
            }
            if j == n || j == i+2 || j-i-2 > 9 {
                return false
            }
            tagName := code[i+2:j]
            if !check(tagName) {
                return false
            }
            if len(stack) == 0 || stack[len(stack)-1] != tagName {
                return false
            }
            stack = stack[:len(stack)-1]
            i = j + 1
        } else if i+1 < n && code[i] == '<' {
            j := i + 1
            for j < n && code[j] != '>' {
                j++
            }
            if j == n || j == i+1 || j-i-1 > 9 {
                return false
            }
            tagName := code[i+1:j]
            if !check(tagName) {
                return false
            }
            stack = append(stack, tagName)
            i = j + 1
        } else {
            i++
        }
    }
    if code == "<" {
        return false
    } else {
        return len(stack) == 0
    }
}

func check(tagName string) bool {
    if len(tagName) == 0 || len(tagName) > 9 {
        return false
    }
    for _, ch := range tagName {
        if ch < 'A' || ch > 'Z' {
            return false
        }
    }
    return true
}
