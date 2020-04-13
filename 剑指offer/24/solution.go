package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	h := head.Next
	head.Next = nil
	t := head

	for h != nil {
		nh := h.Next
		h.Next = t
		t = h
		h = nh
	}
	return t
}
