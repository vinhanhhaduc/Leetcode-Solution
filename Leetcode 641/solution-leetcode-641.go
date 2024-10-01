type MyCircularDeque struct {
    data   []int
    front  int
    rear   int
    size   int
    length int
}
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        data:   make([]int, k+1),
        front:  0,
        rear:   0,
        size:   k + 1,
        length: 0,
    }
}
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.front = (this.front - 1 + this.size) % this.size
    this.data[this.front] = value
    this.length++
    return true
}
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    this.data[this.rear] = value
    this.rear = (this.rear + 1) % this.size
    this.length++
    return true
}
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.front = (this.front + 1) % this.size
    this.length--
    return true
}
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.rear = (this.rear - 1 + this.size) % this.size
    this.length--
    return true
}

func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.front]
}
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[(this.rear-1+this.size)%this.size]
}

func (this *MyCircularDeque) IsEmpty() bool {
    return this.length == 0
}
func (this *MyCircularDeque) IsFull() bool {
    return this.length == this.size-1
}