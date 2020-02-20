package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	q := list.New()
	q.Len()
	if root == nil {
		return []int{}
	}

	q := make([]*TreeNode, 1)
	f, e := -1, 0
	q[e] = root
	output := []int{}

	for e-f > 0 {
		f++
		tn := q[f]

		output = append(output, tn.Val)

		if tn.Left != nil {
			e++
			q = append(q, tn.Left)
		}
		if tn.Right != nil {
			e++
			q = append(q, tn.Right)
		}
	}
	return output
}

func main() {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	tree.Right = &TreeNode{Val: 3}

	fmt.Println(levelOrder(tree))
}
