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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	r := 0
	for r, _ = range inorder {
		if inorder[r] == preorder[0] {
			break
		}
	}

	if r > 0 {
		root.Left = buildTree(preorder[1:r+1], inorder[:r])
	}

	if r < len(inorder) {
		root.Right = buildTree(preorder[r+1:], inorder[r+1:])
	}

	return root
}
