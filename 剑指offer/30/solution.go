package solution

type MinStack struct {
	s []int
	h []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		s: []int{},
		h: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
	if len(this.h) == 0 || this.h[len(this.h)-1] >= x {
		this.h = append(this.h, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.s) == 0 {
		return
	} else {
		p := this.s[len(this.s)-1]
		this.s = this.s[:len(this.s)-1]
		if p == this.h[len(this.h)-1] {
			this.h = this.h[:len(this.h)-1]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.s) == 0 {
		return 0
	}
	return this.s[len(this.s)-1]
}

func (this *MinStack) Min() int {
	if len(this.h) == 0 {
		return 0
	}
	return this.h[len(this.h)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
