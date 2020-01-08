package main

import (
	"fmt"
)

type MyQueue struct {
	stackIn  Stack
	stackOut Stack
}

type Stack struct {
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	x := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return x
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Peek() int {
	return s.data[s.Len()-1]
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stackIn:  Stack{},
		stackOut: Stack{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.stackOut.Len() > 0 {
		return this.stackOut.Pop()
	}

	num := this.stackIn.Len()
	for i := 0; i < num; i++ {
		this.stackOut.Push(this.stackIn.Pop())
	}

	return this.stackOut.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.stackOut.Len() > 0 {
		return this.stackOut.Peek()
	}
	for i := 0; i < this.stackIn.Len(); i++ {
		this.stackOut.Push(this.stackIn.Pop())
	}

	return this.stackOut.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stackOut.Len() == 0 && this.stackIn.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
}
