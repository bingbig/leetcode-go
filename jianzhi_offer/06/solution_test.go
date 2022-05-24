package solution

import (
	"fmt"
	"testing"
)

func TestReversePrint(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	fmt.Println(reversePrint(head))
	fmt.Println(reversePrint(nil))
}

func TestReversePrint2(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	fmt.Println(reversePrint2(head))
	fmt.Println(reversePrint2(nil))
}

// BenchmarkReversePrint-4   	 9485916	       124 ns/op	      56 B/op	       3 allocs/op
func BenchmarkReversePrint(b *testing.B) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	for i := 0; i < b.N; i++ {
		reversePrint(head)
	}
}

// BenchmarkReversePrint2-4   	 9069786	       138 ns/op	      72 B/op	       5 allocs/op
func BenchmarkReversePrint2(b *testing.B) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	for i := 0; i < b.N; i++ {
		reversePrint2(head)
	}
}
