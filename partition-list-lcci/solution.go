package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	p, q := head, head
	for q != nil {
		if q.Val < x {
			q.Val, p.Val = p.Val, q.Val
			p = p.Next
		}
		q = q.Next
	}
	return head
}
