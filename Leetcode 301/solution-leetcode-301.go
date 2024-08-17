func dfs(s string, i, lc, rc, l, r int, path []byte, res *[]string) {
    if i == len(s) {
        if l == 0 && r == 0 {
            *res = append(*res, string(path))
        }
        return
    }

    if s[i] == '(' {
        if l > 0 {
            dfs(s, i+1, lc, rc, l-1, r, path, res)
        }
        path = append(path, '(')
        dfs(s, i+1, lc+1, rc, l, r, path, res)
        path = path[:len(path)-1]
    } else if s[i] == ')' {
        if r > 0 {
            dfs(s, i+1, lc, rc, l, r-1, path, res)
        }
        if lc > rc {
            path = append(path, ')')
            dfs(s, i+1, lc, rc+1, l, r, path, res)
            path = path[:len(path)-1]
        }
    } else {
        path = append(path, s[i])
        dfs(s, i+1, lc, rc, l, r, path, res)
        path = path[:len(path)-1]
    }
}

func removeInvalidParentheses(s string) []string {
    var res []string
    l, r := 0, 0
    for _, ch := range s {
        if ch == '(' {
            l++
        } else if ch == ')' {
            if l > 0 {
                l--
            } else {
                r++
            }
        }
    }

    dfs(s, 0, 0, 0, l, r, []byte{}, &res)
    a := map[string]bool{}
    for _, str := range res {
        a[str] = true
    }
    
    var finalRes []string
    for str := range a {
        finalRes = append(finalRes, str)
    }
    
    return finalRes
}