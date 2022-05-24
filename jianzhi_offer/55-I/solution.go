package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + MaxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
