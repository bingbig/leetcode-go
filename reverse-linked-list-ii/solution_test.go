package solution

import (
	"testing"
)

func TestReverseBetween(t *testing.T) {
	head := &ListNode{
		Val:  3,
		Next: &ListNode{
			Val: 5,
			Next:nil,
		},
	}

	head = ReverseBetween(head, 1, 2)
	t.Log(head.String())
}
