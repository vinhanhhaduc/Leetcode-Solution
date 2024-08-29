type MedianFinder struct {
	left  MaxHeap
	right MinHeap
}
type MaxHeap []int
type MinHeap []int
func Constructor() MedianFinder {
	return MedianFinder{
		left:  MaxHeap{},
		right: MinHeap{},
	}
}
func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 || num <= this.left[0] {
		heap.Push(&this.left, num)
	} else {
		heap.Push(&this.right, num)
	}
	if this.left.Len() > this.right.Len()+1 {
		heap.Push(&this.right, heap.Pop(&this.left))
	} else if this.right.Len() > this.left.Len() {
		heap.Push(&this.left, heap.Pop(&this.right))
	}
}
func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64(this.left[0])
	}
	return float64(this.left[0]+this.right[0]) / 2.0
}
func (h MaxHeap) Len() int            { 
    return len(h) 
}
func (h MaxHeap) Less(i, j int) bool  { 
    return h[i] > h[j] 
}
func (h MaxHeap) Swap(i, j int)       { 
    h[i], h[j] = h[j], h[i]
}
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
func (h MinHeap) Len() int            { 
    return len(h) 
}
func (h MinHeap) Less(i, j int) bool  { 
    return h[i] < h[j] 
}
func (h MinHeap) Swap(i, j int)       { 
    h[i], h[j] = h[j], h[i] 
}
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