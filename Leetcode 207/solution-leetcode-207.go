func canFinish(numCourses int, prerequisites [][]int) bool {
    a := make([][]int, numCourses)
	in := make([]int, numCourses)

	for _, prereq := range prerequisites {
		course := prereq[0]
		req := prereq[1]
		a[req] = append(a[req], course)
		in[course]++
	}
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		count++
		for _, neighbor := range a[curr] {
			in[neighbor]--
			if in[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return count == numCourses
}