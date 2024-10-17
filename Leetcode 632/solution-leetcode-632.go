type Element struct {
    value      int
    listIndex  int
    elemIndex  int
}

type MinHeap []Element

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func smallestRange(nums [][]int) []int {
    h := &MinHeap{}
    heap.Init(h)
    
    maxElement := math.MinInt32
    for i := 0; i < len(nums); i++ {
        heap.Push(h, Element{nums[i][0], i, 0})
        if nums[i][0] > maxElement {
            maxElement = nums[i][0]
        }
    }
    
    res := []int{-1, -1}
    minSize := math.MaxInt32
    for h.Len() == len(nums) {
        minElement := heap.Pop(h).(Element)
        minValue := minElement.value
        listIndex := minElement.listIndex
        elemIndex := minElement.elemIndex
        if maxElement - minValue < minSize {
            minSize = maxElement - minValue
            res[0], res[1] = minValue, maxElement
        }
        if elemIndex + 1 < len(nums[listIndex]) {
            nextValue := nums[listIndex][elemIndex+1]
            heap.Push(h, Element{nextValue, listIndex, elemIndex+1})
            if nextValue > maxElement {
                maxElement = nextValue
            }
        } else {
            break
        }
    }
    
    return res
}