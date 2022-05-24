package solution

import (
	"fmt"
	"testing"
)

var tree = &Node{
	Val: 4,
	Left: &Node{
		Val:   2,
		Left:  &Node{Val: 1},
		Right: &Node{Val: 3},
	},
	Right: &Node{Val: 5},
}

func TestTreeToDoublyList(t *testing.T) {
	head := treeToDoublyList(tree)
	printDoublyList(head)
}

func printDoublyList(head *Node) {
	for head != nil {
		fmt.Printf("%d, left: %v; Right: %v]\n", head.Val, head.Left, head.Right)
		head = head.Right
	}
}
