func robotSim(commands []int, obstacles [][]int) int {
    dirs := [][]int{
        {0, 1},
        {1, 0},
        {0, -1},
        {-1, 0},
    }
    obsSet := make(map[[2]int]struct{})
    for _, obs := range obstacles {
        obsSet[[2]int{obs[0], obs[1]}] = struct{}{}
    }
    x, y, dirIdx := 0, 0, 0
    maxDist := 0
    for _, cmd := range commands {
        if cmd == -2 {
            dirIdx = (dirIdx + 3) % 4
        } else if cmd == -1 {
            dirIdx = (dirIdx + 1) % 4
        } else {
            for i := 0; i < cmd; i++ {
                nx, ny := x + dirs[dirIdx][0], y + dirs[dirIdx][1]
                if _, found := obsSet[[2]int{nx, ny}]; found {
                    break
                }
                x, y = nx, ny
                dist := x*x + y*y
                if dist > maxDist {
                    maxDist = dist
                }
            }
        }
    }
    
    return maxDist
}