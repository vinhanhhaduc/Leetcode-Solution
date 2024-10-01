type MyCalendar struct {
    events [][]int
}

func Constructor() MyCalendar {
    return MyCalendar{events: [][]int{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
    for _, event := range this.events {
        if max(start, event[0]) < min(end, event[1]) {
            return false 
        }
    }
    this.events = append(this.events, []int{start, end})
    return true
}