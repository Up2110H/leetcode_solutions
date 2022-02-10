package queue

type MyCircularDeque struct {
	data []int
	head int
	len  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{make([]int, k), 0, 0}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.head = (len(this.data) + this.head - 1) % len(this.data)
	this.data[this.head] = value
	this.len++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.data[(this.head+this.len)%len(this.data)] = value
	this.len++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % len(this.data)
	this.len--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.len = (this.len - 1) % len(this.data)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if !this.IsEmpty() {
		return this.data[this.head]
	}

	return -1
}

func (this *MyCircularDeque) GetRear() int {
	if !this.IsEmpty() {
		return this.data[(this.head+this.len-1)%len(this.data)]
	}

	return -1
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.len == len(this.data)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
