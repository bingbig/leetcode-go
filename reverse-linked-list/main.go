package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  执行用时 : 100.00%
//  内存消耗 : 10.55%
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	rnext := reverseList(head.Next)
	head.Next = nil
	next.Next = head
	return rnext
}

//  执行用时 : 100.00%
//  内存消耗 : 66.97%
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head.Next
	p2 := head.Next.Next
	head.Next = nil
	p1.Next = head

	for {
		if p2 == nil {
			return p1
		}
		h := p1
		p1 = p2
		p2 = p2.Next
		p1.Next = h
	}
}

func main() {
	l := NewList([]int{1, 2, 3, 4, 5})
	fmt.Println(l.ToSlice())
	l = reverseList(l)
	fmt.Println(l.ToSlice())
	l = reverseList2(l)
	fmt.Println(l.ToSlice())
}
