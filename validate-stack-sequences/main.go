package main

import (
	"fmt"
)

type stack struct {
	data []int
	len  int
	cap  int
}

func validateStackSequences(pushed []int, popped []int) bool {
	pushl := len(pushed)
	popl := len(popped)

	if pushl != popl {
		return false
	}
	if pushl == 0 {
		return true
	}

	s := NewStack(pushl)

	s.Push(pushed[0])
	i := 1
	j := 0

	for i < pushl {
		if top, e := s.Top(); e && top == popped[j] {
			j++
			s.Pop()
		} else {
			s.Push(pushed[i])
			i++
		}
	}

	for j < pushl {
		if top, e := s.Top(); e && top == popped[j] {
			s.Pop()
			j++
		} else {
			return false
		}
	}

	return j == pushl
}

func NewStack(length int) *stack {
	return &stack{
		data: make([]int, length),
		len:  0,
		cap:  length,
	}
}

func (s *stack) Top() (int, bool) {
	if s.len == 0 {
		return -1, false
	}

	return s.data[s.len-1], true
}

func (s *stack) Pop() (int, bool) {
	if s.len == 0 {
		return 0, false
	}

	n := s.data[s.len-1]
	s.len--
	return n, true
}

func (s *stack) Push(n int) bool {
	if s.len == s.cap {
		return false
	}

	s.len++
	s.data[s.len-1] = n
	return true
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
