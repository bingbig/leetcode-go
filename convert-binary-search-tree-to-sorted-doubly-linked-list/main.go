package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	stack := list.New()
	p := root
	for p != nil {
		stack.PushBack(p)
		p = p.Left
	}

	head := stack.Back().Value.(*Node)
	stack.Remove(stack.Back())
	if head.Right != nil {
		stack.PushBack(head.Right)
	}

	p = head
	for stack.Len() > 0 {
		next := stack.Back().Value.(*Node)
		stack.Remove(stack.Back())
		if next.Right != nil {
			stack.PushBack(next.Right)
		}
		p.Right = next
		next.Left = p

		p = p.Right
	}

	return head
}

func printDoubleList(head *Node) {
	p := head
	for p != nil {
		fmt.Printf(" %d ", p.Val)
		p = p.Right
	}

	fmt.Println("")
}

func main() {
	root := &Node{
		Val: 4,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 1},
			Right: &Node{Val: 3},
		},
		Right: &Node{Val: 5},
	}

	dlist := treeToDoublyList(root)
	printDoubleList(dlist)

	dlist = treeToDoublyList(nil)
	printDoubleList(dlist)

	dlist = treeToDoublyList(&Node{Val: 1})
	printDoubleList(dlist)
}
