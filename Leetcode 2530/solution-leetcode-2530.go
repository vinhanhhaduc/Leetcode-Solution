type MaxHeap []int
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }  
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func maxKelements(nums []int, k int) int64 {
    h := &MaxHeap{}
    heap.Init(h)
    for _, num := range nums {
        heap.Push(h, num)
    }
    
    var res int64 = 0
    for i := 0; i < k; i++ {
        largest := heap.Pop(h).(int)
        res += int64(largest)
        reduced := (largest + 2) / 3
        heap.Push(h, reduced)
    }
    
    return res
}