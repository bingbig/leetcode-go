package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.

 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	rootNode := &TreeNode{Val: preorder[0]}
	rootInorder := 0

	for idx, n := range inorder {
		if n == rootNode.Val {
			rootInorder = idx
		}
	}

	if rootInorder == 0 {
		rootNode.Right = buildTree(preorder[1:], inorder[1:])
	} else if rootInorder == len(inorder)-1 {
		rootNode.Left = buildTree(preorder[1:], inorder[:rootInorder])
	} else {
		rootNode.Left = buildTree(preorder[1:rootInorder+1], inorder[:rootInorder])
		rootNode.Right = buildTree(preorder[rootInorder+1:], inorder[rootInorder+1:])
	}

	return rootNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	fmt.Println(tree.PreorderString())
}
