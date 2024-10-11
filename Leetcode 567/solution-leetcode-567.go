func checkInclusion(s1 string, s2 string) bool {
    n1, n2 := len(s1), len(s2)
    if n1 > n2 {
        return false
    }
    s1Freq := make([]int, 26)
    s2Freq := make([]int, 26)
    for i := 0; i < n1; i++ {
        s1Freq[s1[i]-'a']++
        s2Freq[s2[i]-'a']++
    }
    for i := 0; i <= n2-n1; i++ {
        if matches(s1Freq, s2Freq) {
            return true
        }
        if i+n1 < n2 {
            s2Freq[s2[i]-'a']--   
            s2Freq[s2[i+n1]-'a']++  
        }
    }

    return false
}
func matches(arr1, arr2 []int) bool {
    for i := 0; i < 26; i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}
