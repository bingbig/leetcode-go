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
	isOdd := 0
	output := [][]int{}
	for queue.Len() > 0 {
		nodes := queue.Len()
		tmp := make([]int, nodes)
		for i := 0; i < nodes; i++ {
			tn := queue.Front().Value.(*TreeNode)
			idx := i
			if isOdd&1 == 1 {
				idx = nodes - 1 - i
			}
			tmp[idx] = tn.Val
			if tn.Left != nil {
				queue.PushBack(tn.Left)
			}
			if tn.Right != nil {
				queue.PushBack(tn.Right)
			}
			queue.Remove(queue.Front())
		}
		output = append(output, tmp)
		isOdd++
	}
	return output
}

func main() {
	tree := &TreeNode{Val: 3}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}

	fmt.Println(levelOrder(tree))
}
