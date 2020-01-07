package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tree *TreeNode) PreorderString() string {
	if tree == nil {
		return ""
	}

	if tree.Left == nil && tree.Right == nil {
		return fmt.Sprintf("%d\t", tree.Val)
	}

	str := fmt.Sprintf("%d\t", tree.Val)

	if tree.Left == nil {
		str += "nil\t"
	} else {
		str += tree.Left.PreorderString()
	}

	if tree.Right == nil {
		str += "nil\t"
	} else {
		str += tree.Right.PreorderString()
	}

	return str
}

func (tree *TreeNode) InorderString() string {
	if tree == nil {
		return ""
	}

	if tree.Left == nil && tree.Right == nil {
		return fmt.Sprintf("%d\t", tree.Val)
	}

	str := fmt.Sprintf("%d\t", tree.Val)
	if tree.Left != nil {
		str = tree.Left.InorderString() + str
	}
	if tree.Right != nil {
		str += tree.Right.InorderString()
	}

	return str
}
