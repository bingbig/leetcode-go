package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int, pos int) *ListNode {
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

	if pos >= 0 {
		joint := l
		for i := 0; i < pos; i++ {
			joint = joint.Next
		}

		h.Next = joint
	}

	return l
}
