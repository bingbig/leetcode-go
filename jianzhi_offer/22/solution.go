package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k <= 0 {
		return nil
	}

	p1, p2 := head, head
	for i := 1; i <= k; i++ {
		p1 = p1.Next
		if i != k && p1 == nil {
			return nil
		}
	}

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}
