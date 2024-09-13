func numBusesToDestination(routes [][]int, source int, target int) int {
    if source == target {
		return 0
	}
	s := make(map[int][]int)
	for routeIndex, route := range routes {
		for _, stop := range route {
			s[stop] = append(s[stop], routeIndex)
		}
	}
	queue := list.New()
	visitedRoutes := make(map[int]bool)
	visitedStops := make(map[int]bool)
	for _, routeIndex := range s[source] {
		queue.PushBack(routeIndex)
		visitedRoutes[routeIndex] = true
	}
	visitedStops[source] = true
	count := 1
	for queue.Len() > 0 {
		queueSize := queue.Len()

		for i := 0; i < queueSize; i++ {
			currentRoute := queue.Remove(queue.Front()).(int)
			for _, stop := range routes[currentRoute] {
				if stop == target {
					return count
				}
				if !visitedStops[stop] {
					for _, nextRoute := range s[stop] {
						if !visitedRoutes[nextRoute] {
							queue.PushBack(nextRoute)
							visitedRoutes[nextRoute] = true
						}
					}
					visitedStops[stop] = true
				}
			}
		}
		count++
	}
	return -1
}