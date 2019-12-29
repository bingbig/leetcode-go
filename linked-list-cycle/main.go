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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	p1 := head
	p2 := head.Next
	for {
		if p2 == nil || p2.Next == nil {
			return false
		}
		if p1 == p2 || p1 == p2.Next {
			return true
		}

		p1 = p1.Next
		p2 = p2.Next.Next
	}
}

func main() {
	fmt.Println(hasCycle(NewList([]int{3, 2, 0, -4}, 1)))
	fmt.Println(hasCycle(NewList([]int{3, 2}, 0)))
	fmt.Println(hasCycle(NewList([]int{3}, -1)))
}
