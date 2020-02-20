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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	queue.PushBack(root)
	output := [][]int{}
	for queue.Len() > 0 {
		nodes := queue.Len()
		tmp := make([]int, nodes)
		for i := 0; i < nodes; i++ {
			tn := queue.Front().Value.(*TreeNode)
			tmp[i] = tn.Val
			if tn.Left != nil {
				queue.PushBack(tn.Left)
			}
			if tn.Right != nil {
				queue.PushBack(tn.Right)
			}
			queue.Remove(queue.Front())
		}
		output = append(output, tmp)
	}
	return output
}

func main() {
	tree := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(levelOrder(tree))
}
