func findChampion(n int, edges [][]int) int {
    check := make([]bool, n)
    for i := range check {
        check[i] = true
    }
    
    for _, edge := range edges {
        _, loser := edge[0], edge[1]
        check[loser] = false
    }
    
    champion := -1
    res := 0
    
    for team := 0; team < n; team++ {
        if check[team] {
            champion = team
            res++
        }
    }
    
    if res == 1 {
        return champion
    }
    return -1
}