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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head

	for {
		if fast == nil {
			tail := reverseList(slow)
			return containList(head, tail)
		}
		if fast.Next == nil {
			tail := reverseList(slow.Next)
			return containList(head, tail)
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
}

func containList(l1, l2 *ListNode) bool {
	for {
		if l2 == nil {
			return true
		} else if l1 == nil {
			return false
		} else {
			if l1.Val != l2.Val {
				return false
			} else {
				l1 = l1.Next
				l2 = l2.Next
			}
		}
	}
}

func reverseList(head *ListNode) *ListNode {
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
	fmt.Println(isPalindrome(NewList([]int{1, 2})))
	fmt.Println(isPalindrome(NewList([]int{1, 2, 2, 1})))
	fmt.Println(isPalindrome(NewList([]int{1})))
	fmt.Println(isPalindrome(NewList([]int{})))
	fmt.Println(isPalindrome(NewList([]int{1, 2, 3, 2, 1})))
	fmt.Println(isPalindrome(NewList([]int{1, 2, 3, 3, 2, 1, 0})))
}
