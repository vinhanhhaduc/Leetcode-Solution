func reverseParentheses(s string) string {
    str := []string{}
    for _, i := range s {
       if i == ')' {
            temp := []string{}
            for len(str) > 0 {
                k := str[len(str) - 1]
                str = str[:len(str) - 1]
                if k == "(" {
                    break
                }
                temp = append(temp, k)
                
            }
            str = append(str, temp...)
       } else {
            str = append(str, string(i));
       }
    }
    res := strings.Join(str, "")
    return res
}
