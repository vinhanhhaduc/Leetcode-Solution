func maximumGain(s string, x int, y int) int {
    val, value := byte('a'), byte('b')
    if x < y {
        val, value = value, val
        x, y = y, x
    }

    res := 0
    cnt, count := 0, 0
    for i := range s {
        if s[i] == val {
            cnt++
            continue
        }
        if s[i] == value {
            if cnt > 0 {
                cnt--
                res += x
            } else {
                count++
            }
        } else {
            res += y * min(cnt, count)
            cnt, count = 0, 0
        }
    }

    res += y * min(cnt, count)

    return res
}

