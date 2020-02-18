// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
package main

import (
	"fmt"
)

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}

	if A.Val == B.Val {
		if cmpTrees(A, B) {
			return true
		}
	}

	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func cmpTrees(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil {
		return false
	}

	if A.Val != B.Val {
		return false
	} else {
		return cmpTrees(A.Left, B.Left) && cmpTrees(A.Right, B.Right)
	}
}

func main() {
	treeA := NewTreeNode(3)
	treeA.Left = NewTreeNode(4)
	treeA.Left.Left = NewTreeNode(1)
	treeA.Left.Right = NewTreeNode(2)
	treeA.Right = NewTreeNode(5)

	treeB := NewTreeNode(4)
	treeB.Left = NewTreeNode(1)

	fmt.Println(isSubStructure(treeA, treeB))
}
