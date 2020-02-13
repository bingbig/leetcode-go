package main

import (
	"fmt"
)

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k == 0 {
		return nil
	}

	n1 := head
	n2 := head
	for i := 0; i < k; i++ {
		if n1 == nil {
			return nil
		}
		n1 = n1.Next
	}

	for n1 != nil {
		n1 = n1.Next
		n2 = n2.Next
	}

	return n2
}

func main() {
	list := NewList([]int{1, 2, 3, 4, 5, 6, 7, 9, 10})
	fmt.Println(getKthFromEnd(list, 3))
}
