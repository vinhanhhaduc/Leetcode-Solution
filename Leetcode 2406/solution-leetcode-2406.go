func minGroups(intervals [][]int) int {
     events := []struct {
        time int
        typeEvent int
    }{}
    for _, interval := range intervals {
        events = append(events, struct{time, typeEvent int}{interval[0], 1})
        events = append(events, struct{time, typeEvent int}{interval[1] + 1, -1}) 
    }
    sort.Slice(events, func(i, j int) bool {
        if events[i].time == events[j].time {
            return events[i].typeEvent < events[j].typeEvent
        }
        return events[i].time < events[j].time
    })

    res := 0
    active := 0
    for _, event := range events {
        active += event.typeEvent
        if active > res {
            res = active
        }
    }

    return res
}