package reverse_linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	head = ReverseList(head)
	assert.Equal(t, "head->5->4->3->2->1", ReverseList(head))
}

func (l *ListNode) String() string {
	h := l
	output := "head"
	for h != nil {
		output = fmt.Sprintf("%s->%d", output, h.Val)
		h = h.Next
	}
	return output
}
