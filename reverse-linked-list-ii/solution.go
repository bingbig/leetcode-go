package solution

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	length := right - left

	var pv, p, preLinkTail, nextLinkHead, innerLinkHead *ListNode
	p = head
	for left > 1 {
		left--
		pv = p
		p = p.Next
	}
	preLinkTail, innerLinkHead = pv, p
	p = innerLinkHead
	for length > 0 {
		length--
		p = p.Next
	}
	nextLinkHead = p.Next
	p.Next = nil

	innerLinkHead = ReverseList(innerLinkHead)
	if preLinkTail != nil {
		preLinkTail.Next = innerLinkHead
	} else {
		head = innerLinkHead
	}
	p = innerLinkHead
	for {
		if p.Next == nil {
			break
		}
		p = p.Next
	}
	p.Next = nextLinkHead

	return head
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, curr, next *ListNode
	curr, next = head, head.Next
	for next != nil {
		curr.Next = prev
		prev, curr, next = curr, next, next.Next
	}
	curr.Next = prev
	return curr
}

func (l *ListNode) String() string {
	h := l
	output := "head"
	for h != nil {
		output = fmt.Sprintf("%s->%d", output, h.Val)
		h = h.Next
	}
	return output
}
