package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
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
