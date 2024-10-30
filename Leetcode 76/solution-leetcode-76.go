func minWindow(s string, t string) string {
     if len(t) > len(s) {
        return ""
    }

    cnt := make(map[byte]int)
    for i := range t {
        cnt[t[i]]++
    }

    wc := make(map[byte]int)
    have, need := 0, len(cnt)
    left := 0
    minLen := len(s) + 1
    start := 0

    for right := 0; right < len(s); right++ {
        char := s[right]
        wc[char]++
        
        if count, ok := cnt[char]; ok && wc[char] == count {
            have++
        }

        for have == need {
            if right - left + 1 < minLen {
                minLen = right - left + 1
                start = left
            }
            
            wc[s[left]]--
            if count, ok := cnt[s[left]]; ok && wc[s[left]] < count {
                have--
            }
            left++
        }
    }

    if minLen == len(s) + 1 {
        return ""
    }
    return s[start : start + minLen]
}