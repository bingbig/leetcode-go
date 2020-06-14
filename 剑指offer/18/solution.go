package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return head.Next
	}

	p1 := head.Next
	if p1.Val == val {
		head.Next = p1.Next
		return head
	}

	p2 := p1.Next

	for p2 != nil {
		if p2.Val == val {
			p1.Next = p2.Next
			return head
		}
		p1 = p2
		p2 = p2.Next
	}

	return head
}
