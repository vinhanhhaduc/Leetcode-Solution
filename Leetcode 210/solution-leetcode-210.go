func findOrder(numCourses int, prerequisites [][]int) []int {
    a := make([][]int, numCourses)
	id := make([]int, numCourses)
	for _, prereq := range prerequisites {
		course, prereqCourse := prereq[0], prereq[1]
		a[prereqCourse] = append(a[prereqCourse], course)
		id[course]++
	}
	queue := []int{}
	for course := 0; course < numCourses; course++ {
		if id[course] == 0 {
			queue = append(queue, course)
		}
	}
	var result []int

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)
		for _, neighbor := range a[current] {
			id[neighbor]--
			if id[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(result) == numCourses {
		return result
	}
	return []int{}
}