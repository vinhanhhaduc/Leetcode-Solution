type MyCalendarTwo struct {
    single []Interval
    doubleBooked []Interval
}

type Interval struct {
    start, end int
}

func Constructor() MyCalendarTwo {
    return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
    for _, booking := range this.doubleBooked {
        if max(start, booking.start) < min(end, booking.end) {
            return false
        }
    }
    for _, booking := range this.single {
        if max(start, booking.start) < min(end, booking.end) {
            this.doubleBooked = append(this.doubleBooked, Interval{max(start, booking.start), min(end, booking.end)})
        }
    }
    this.single = append(this.single, Interval{start, end})
    return true
}