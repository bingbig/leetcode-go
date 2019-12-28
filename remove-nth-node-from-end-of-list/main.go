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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	if head.Next == nil && n == 1 {
		return nil
	}

	p1 := head
	p2 := head
	for i := 1; i <= n; i++ {
		p1 = p1.Next
	}

	if p1 == nil {
		return head.Next
	}

	for {
		if p1.Next == nil {
			p2.Next = p2.Next.Next
			return head
		}
		p1 = p1.Next
		p2 = p2.Next
	}
}

func main() {
	l := NewList([]int{1, 2})
	l = removeNthFromEnd(l, 2)
	fmt.Println(l.ToSlice())
}
