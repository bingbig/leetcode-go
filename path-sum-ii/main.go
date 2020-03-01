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

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil || root.Val > sum {
		return [][]int{}
	}

	if root.Val == sum {
		if root.Left == nil || root.Right == nil {
			return [][]int{[]int{root.Val}}
		} else {
			return [][]int{}
		}
	}

	stack := list.New()
	if root.Right != nil {
		stack.PushBack(root.Right)
	}

	paths := [][]int{}
	path := []int{root.Val}
	expect := sum - root.Val
	tn := root.Left
	for stack.Len() > 0 || tn != nil {
		fmt.Println(stack, expect, path)
		if tn == nil {
			if expect == 0 {
				save := make([]int, len(path))
				copy(save, path)
				paths = append(paths, save)
				tn = stack.Back().Value.(*TreeNode)
				stack.Remove(stack.Back())
				expect = expect + path[len(path)-1]
				path = path[:len(path)-2]
			} else {
				tn = stack.Back().Value.(*TreeNode)
				stack.Remove(stack.Back())
			}
		} else {
			if tn.Right != nil {
				stack.PushBack(tn.Right)
			}

			if expect < tn.Val {
				tn = stack.Back().Value.(*TreeNode)
				stack.Remove(stack.Back())
			} else {
				expect = expect - tn.Val
				path = append(path, tn.Val)
				if expect == 0 && tn.Left == nil && tn.Right == nil {
					save := make([]int, len(path))
					copy(save, path)
					paths = append(paths, save)
				}
				tn = tn.Left
			}
		}

	}

	return paths
}

// TODO: solution incomplete
func main() {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	// tree2 := &TreeNode{
	// 	Val: 5,
	// 	Right: &TreeNode{
	// 		Val: 8,
	// 		Right: &TreeNode{
	// 			Val:  4,
	// 			Left: &TreeNode{Val: 5},
	// 		},
	// 	},
	// }

	fmt.Println(pathSum(tree, 22))
	// fmt.Println(pathSum(tree2, 22))
}
