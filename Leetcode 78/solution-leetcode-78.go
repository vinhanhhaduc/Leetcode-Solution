func subsets(nums []int) [][]int {
    res := [][]int{{}}

    for _, num := range nums {
        var n [][]int
        for _, subset := range res {
            ns := append([]int(nil), subset...)
            ns = append(ns, num)
            n = append(n, ns)
        }
        res = append(res, n...)
    }

    return res
}