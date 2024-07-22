type Person struct {
	name   string
	height int
}
func sortPeople(names []string, heights []int) []string {
    n := len(names)
	people := make([]Person, n)

	for i := 0; i < n; i++ {
		people[i] = Person{name: names[i], height: heights[i]}
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].height > people[j].height
	})

	sortedNames := make([]string, n)
	for i := 0; i < n; i++ {
		sortedNames[i] = people[i].name
	}
	return sortedNames

}