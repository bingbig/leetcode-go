package solution

type MaxQueue struct {
	Q []int
	H []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		Q: []int{},
		H: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.H) == 0 {
		return -1
	}

	return this.H[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.Q = append(this.Q, value)

	if len(this.H) > 0 && this.H[0] < value {
		this.H = []int{value}
	} else {
		for len(this.H) > 0 && this.H[len(this.H)-1] < value {
			this.H = this.H[:len(this.H)-1]
		}
		this.H = append(this.H, value)
	}
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Q) == 0 {
		return -1
	} else if len(this.Q) == 1 {
		f := this.Q[0]
		this.Q = []int{}
		this.H = []int{}
		return f
	} else {
		f := this.Q[0]
		this.Q = this.Q[1:]
		if this.H[0] == f {
			this.H = this.H[1:]
		}
		return f
	}
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
