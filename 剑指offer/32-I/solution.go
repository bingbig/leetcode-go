package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	output := make([]int, 0)
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		output = append(output, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return output
}
