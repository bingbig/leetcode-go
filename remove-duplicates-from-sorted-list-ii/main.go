package main

import (
	"fmt"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var newHead, newTail *ListNode
	prev := &ListNode{Val: head.Val - 1, Next: head}
	cur := head
	next := head.Next

	for next != nil {
		if cur.Val != prev.Val && cur.Val != next.Val {
			if newHead == nil {
				newHead = &ListNode{Val: cur.Val}
				newTail = newHead
			} else {
				newTail.Next = &ListNode{Val: cur.Val}
				newTail = newTail.Next
			}
		}

		prev = prev.Next
		cur = cur.Next
		next = next.Next
	}

	if cur.Val != prev.Val {
		if newHead == nil {
			newHead = &ListNode{Val: cur.Val}
		} else {
			newTail.Next = &ListNode{Val: cur.Val}
		}
	}

	return newHead
}

func main() {
	list := NewList([]int{1, 1, 2})
	fmt.Println(list.ToSlice())
	fmt.Println(deleteDuplicates(list).ToSlice())
}
