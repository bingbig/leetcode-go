package solution

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: -9, Next: &ListNode{Val: -2}}
	l2 := &ListNode{Val: 3, Next: &ListNode{Val: 6}}

	fmt.Println(mergeTwoLists(l1, l2))
}
