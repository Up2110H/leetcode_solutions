package stack

type customNode struct {
	val  int
	next *customNode
}

type CustomStack struct {
	top     *customNode
	size    int
	maxSize int
}

func CustomStackConstructor(maxSize int) CustomStack {
	return CustomStack{maxSize: maxSize}
}

func (this *CustomStack) Push(x int) {
	if this.size < this.maxSize {
		this.top = &customNode{x, this.top}
		this.size++
	}
}

func (this *CustomStack) Pop() int {
	if this.size == 0 {
		return -1
	}

	this.size--
	defer func() {
		this.top = this.top.next
	}()
	return this.top.val
}

func (this *CustomStack) Increment(k int, val int) {
	it := this.top
	for i := this.size; i > 0; i-- {
		if i <= k {
			it.val += val
		}
		it = it.next
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
