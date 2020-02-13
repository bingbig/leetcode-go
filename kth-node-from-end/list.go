package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(v int) *ListNode {
	return &ListNode{Val: v}
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
