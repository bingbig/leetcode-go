package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	l := postorderTraversal2(root.Left)
	r := postorderTraversal2(root.Right)

	return append(append(l, r...), root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	m := len(res) - 1
	for i := 0; i < m/2+1; i++ {
		res[i], res[m-i] = res[m-i], res[i]
	}

	return res
}
