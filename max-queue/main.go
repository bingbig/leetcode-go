package main

import "fmt"

type MaxQueue struct {
	M []int
	Q []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		M: make([]int, 0),
		Q: make([]int, 0),
	}
}

func (this *MaxQueue) Is_Empty() bool {
	return len(this.Q) == 0
}

func (this *MaxQueue) Max_value() int {
	if this.Is_Empty() {
		return -1
	}
	return this.M[0]
}

func (this *MaxQueue) Push_back(value int) {
	for i := len(this.M) - 1; !this.Is_Empty() && i >= 0; i-- {
		if this.M[i] < value {
			this.M = this.M[:i]
		}
	}
	this.M = append(this.M, value)
	this.Q = append(this.Q, value)
}

func (this *MaxQueue) Pop_front() int {
	if this.Is_Empty() {
		return -1
	}

	if this.Q[0] == this.M[0] {
		this.M = this.M[1:]
	}
	front := this.Q[0]
	this.Q = this.Q[1:]
	return front
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

func main() {
	mq := Constructor()
	fmt.Println(mq.Is_Empty())

	mq.Push_back(1)
	fmt.Println(mq.Max_value())
	mq.Push_back(100)
	fmt.Println(mq.Max_value())
	fmt.Println(mq.Pop_front())
	fmt.Println(mq.Max_value())
}
