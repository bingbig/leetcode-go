package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	oq, eq := make([]*TreeNode, 1), []*TreeNode{}
	oq[0] = root
	for len(oq) > 0 {
		level := make([]int, len(oq))
		for i := 0; i < len(oq); i++ {
			level[i] = oq[i].Val
			if oq[i].Left != nil {
				eq = append(eq, oq[i].Left)
			}
			if oq[i].Right != nil {
				eq = append(eq, oq[i].Right)
			}
		}
		oq, eq = eq, []*TreeNode{}
		res = append(res, level)
	}

	return res
}
