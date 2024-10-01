type CustomStack struct {
    stack      []int
    incTracker []int
    maxSize    int  
}
func Constructor(maxSize int) CustomStack {
    return CustomStack{
        stack:      []int{},
        incTracker: make([]int, maxSize),
        maxSize:    maxSize,
    }
}
func (this *CustomStack) Push(x int) {
    if len(this.stack) < this.maxSize {
        this.stack = append(this.stack, x)
    }
}
func (this *CustomStack) Pop() int {
    if len(this.stack) == 0 {
        return -1
    }

    index := len(this.stack) - 1
    res := this.stack[index] + this.incTracker[index]

    if index > 0 {
        this.incTracker[index-1] += this.incTracker[index]
    }
    this.incTracker[index] = 0
    this.stack = this.stack[:index]

    return res
}
func (this *CustomStack) Increment(k int, val int) {
    if len(this.stack) == 0 {
        return
    }
    limit := min(k, len(this.stack))
    this.incTracker[limit-1] += val
}