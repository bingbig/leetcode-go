package main

import (
	"fmt"
)

type MinStack struct {
	Helper []int
	Data   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Data:   make([]int, 0),
		Helper: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if len(this.Helper) > 0 {
		if this.Helper[len(this.Helper)-1] >= x {
			this.Helper = append(this.Helper, x)
		}
	} else {
		this.Helper = append(this.Helper, x)
	}
}

func (this *MinStack) Pop() {
	length := len(this.Data)
	if length == 0 {
		return
	} else if length == 1 {
		this.Data = []int{}
		this.Helper = []int{}
	} else {
		t := this.Data[length-1]
		this.Data = this.Data[:length-1]
		hl := len(this.Helper)
		if hl > 0 {
			if t == this.Helper[hl-1] {
				this.Helper = this.Helper[:len(this.Helper)-1]
			}
		}
	}
}

func (this *MinStack) Top() int {
	dl := len(this.Data)
	if dl == 0 {
		return 0
	} else {
		return this.Data[dl-1]
	}
}

func (this *MinStack) GetMin() int {
	if len(this.Helper) > 0 {
		return this.Helper[len(this.Helper)-1]
	} else {
		return 0
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
