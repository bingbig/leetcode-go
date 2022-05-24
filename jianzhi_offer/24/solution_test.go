package solution

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	rlist := reverseList(list)
	rrlist := reverseList(rlist)
	if !reflect.DeepEqual(list, rrlist) {
		t.Error("failed")
	}
	fmt.Printf("%+v", reverseList(rlist))
}
