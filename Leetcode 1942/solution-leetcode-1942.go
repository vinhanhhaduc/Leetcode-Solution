import (
    "container/heap"
    "sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type Event struct {
    time   int
    person int
    leaving bool
}

func smallestChair(times [][]int, targetFriend int) int {
    n := len(times)
    events := make([]Event, 0, 2*n)
    for i, t := range times {
        events = append(events, Event{t[0], i, false}) 
        events = append(events, Event{t[1], i, true})
    }

    sort.Slice(events, func(i, j int) bool {
        if events[i].time == events[j].time {
            return events[i].leaving
        }
        return events[i].time < events[j].time
    })

    availableChairs := &MinHeap{}
    heap.Init(availableChairs)
    for i := 0; i < n; i++ {
        heap.Push(availableChairs, i)
    }

    occupiedChairs := make(map[int]int)

    for _, event := range events {
        if event.leaving {
            chair := occupiedChairs[event.person]
            heap.Push(availableChairs, chair)
        } else {
            chair := heap.Pop(availableChairs).(int)
            occupiedChairs[event.person] = chair

            if event.person == targetFriend {
                return chair
            }
        }
    }

    return -1
}
