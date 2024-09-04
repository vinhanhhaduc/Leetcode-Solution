func findRedundantDirectedConnection(edges [][]int) []int {
    n := len(edges)
    parent := make([]int, n+1)
    for i := range parent {
        parent[i] = i
    }
    
    un := make([]int, n+1)
    for i := range un {
        un[i] = i
    }

    var ce, cye []int
    
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        
        if parent[v] != v {
            ce = edge
        } else {
            parent[v] = u
            
            if helper(un, u) == helper(un, v) {
                cye = edge
            } else {
                helper2(un, u, v)
            }
        }
    }

    if ce == nil {
        return cye
    }
    
    if cye == nil {
        return ce
    }
    
    for _, edge := range edges {
        if edge[1] == ce[1] {
            return edge
        }
    }
    
    return []int{}
}

func helper(parent []int, i int) int {
    for parent[i] != i {
        parent[i] = parent[parent[i]]
        i = parent[i]
    }
    return i
}

func helper2(parent []int, i, j int) {
    ai := helper(parent, i)
    aj := helper(parent, j)
    if ai != aj {
        parent[ai] = aj
    }
}
