package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isb, _ := isBalancedWithDept(root)
	return isb
}

func isBalancedWithDept(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	isleftBalanced, left := isBalancedWithDept(root.Left)
	isRightBalanced, right := isBalancedWithDept(root.Right)
	if isleftBalanced && isRightBalanced {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			return true, maxInt(left, right) + 1
		}
	}
	return false, 0
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
