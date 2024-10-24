func parseBoolExpr(expression string) bool {
    stk := []rune{}
    for _, c := range expression {
        if c != ')' && c != ',' {
            stk = append(stk, c)
        } else if c == ')' {
            exp := []bool{}
            for len(stk) > 0 && stk[len(stk)-1] != '(' {
                t := stk[len(stk)-1]
                stk = stk[:len(stk)-1]
                exp = append(exp, t == 't')
            }
            
            stk = stk[:len(stk)-1]
            
            if len(stk) > 0 {
                t := stk[len(stk)-1]
                stk = stk[:len(stk)-1]
                v := exp[0]
                if t == '&' {
                    for _, b := range exp {
                        v = v && b
                    }
                } else if t == '|' {
                    for _, b := range exp {
                        v = v || b
                    }
                } else {  // NOT operation
                    v = !v
                }
                if v {
                    stk = append(stk, 't')
                } else {
                    stk = append(stk, 'f')
                }
            }
        }
    }
    return stk[len(stk)-1] == 't'
}