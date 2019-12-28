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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := l1
	if l1.Val < l2.Val {
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	p := head

	for {
		if l1 == nil && l2 == nil {
			return head
		}

		if l1 == nil {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
			continue
		}
		if l2 == nil {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
			continue
		}
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
}

func main() {
	l1 := NewList([]int{1, 2, 4})
	l2 := NewList([]int{1, 3, 4})
	ml := mergeTwoLists(l1, l2)
	fmt.Println(ml.ToSlice())
}
