func outerTrees(trees [][]int) [][]int {
    helper := func(p, q, r []int) int {
        return (q[0]-p[0])*(r[1]-p[1]) - (q[1]-p[1])*(r[0]-p[0])
    }

    distance := func(p, q []int) int {
        return (q[0]-p[0])*(q[0]-p[0]) + (q[1]-p[1])*(q[1]-p[1])
    }

    if len(trees) <= 3 {
        return trees
    }
    start := 0
    for i := 1; i < len(trees); i++ {
        if trees[i][0] < trees[start][0] || (trees[i][0] == trees[start][0] && trees[i][1] < trees[start][1]) {
            start = i
        }
    }

    res := [][]int{}
    visited := make([]bool, len(trees))
    curr := start

    for {
        res = append(res, trees[curr])
        visited[curr] = true
        next := (curr + 1) % len(trees)

        for i := 0; i < len(trees); i++ {
            cross := helper(trees[curr], trees[i], trees[next])
            if cross > 0 || (cross == 0 && distance(trees[curr], trees[i]) > distance(trees[curr], trees[next])) {
                next = i
            }
        }

        for i := 0; i < len(trees); i++ {
            if i != curr && i != next && helper(trees[curr], trees[next], trees[i]) == 0 {
                if !visited[i] {
                    res = append(res, trees[i])
                    visited[i] = true
                }
            }
        }
        curr = next
        if curr == start {
            break
        }
    }

    return res
}