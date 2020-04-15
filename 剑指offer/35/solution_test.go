package solution

import (
	"fmt"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	node4 := &Node{Val: 1}
	node3 := &Node{Val: 10, Next: node4}
	node2 := &Node{Val: 11, Next: node3}
	node1 := &Node{Val: 13, Next: node2}
	node0 := &Node{Val: 7, Next: node1}
	node1.Random = node0
	node2.Random = node4
	node3.Random = node2
	node4.Random = node0
	list := node0
	printList(list)
	fmt.Println()
	nlist := copyRandomList(list)
	printList(list)
	fmt.Println()
	printList(nlist)
}

func printList(head *Node) {
	for head != nil {
		rand := -1
		if head.Random != nil {
			rand = head.Random.Val
		}

		fmt.Println([]int{head.Val, rand})
		head = head.Next
	}
}
