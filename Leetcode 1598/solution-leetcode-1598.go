func minOperations(logs []string) int {
    cnt := 0
    for _, i := range logs {
        if i == "../" {
            if cnt > 0 {
                cnt--
            }
        } else if i == "./" {
            continue
        } else {
            cnt++
        }
    }
    return cnt
}