package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) *ListNode {
	c := len(nums)
	if c == 0 {
		return nil
	}
	l := &ListNode{
		Val: nums[0],
	}
	h := l
	for i := 1; i < c; i++ {
		h.Next = &ListNode{
			Val: nums[i],
		}
		h = h.Next
	}

	return l
}

func (l *ListNode) ToSlice() []int {
	if l == nil {
		return []int{}
	}
	s := []int{}
	h := l
	for {
		s = append(s, h.Val)
		if h.Next != nil {
			h = h.Next
		} else {
			return s
		}
	}
}
