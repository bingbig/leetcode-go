package solution

type CustomStack struct {
	idx  int
	cap  int
	data []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		data: make([]int, maxSize),
		idx:  -1,
		cap:  maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if this.idx < this.cap-1 {
		this.idx++
		this.data[this.idx] = x
	}
}

func (this *CustomStack) Pop() int {
	if this.idx < 0 {
		return -1
	}
	v := this.data[this.idx]
	this.idx--
	return v
}

func (this *CustomStack) Increment(k int, val int) {
	if k > this.idx+1 {
		k = this.idx + 1
	}

	for i := 0; i < k; i++ {
		this.data[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
