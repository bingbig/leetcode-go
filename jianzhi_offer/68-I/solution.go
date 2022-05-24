package solution

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 执行用时 :24 ms, 在所有 Go 提交中击败了83.85%的用户
// 内存消耗 :6.8 MB, 在所有 Go 提交中击败了52.63%的用户
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	if q.Val > p.Val {
		q, p = p, q
	}

	if root.Val > q.Val && root.Val < p.Val {
		return root
	}

	if root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if root.Val > p.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return nil
}

// 执行用时 :24 ms, 在所有 Go 提交中击败了83.85%的用户
// 内存消耗 :6.6 MB, 在所有 Go 提交中击败了94.74%的用户
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	if q.Val > p.Val {
		q, p = p, q
	}

	ptr := root
	for ptr != nil {
		if ptr.Val >= q.Val && ptr.Val <= p.Val {
			return ptr
		} else if ptr.Val < q.Val {
			ptr = ptr.Right
		} else if ptr.Val > p.Val {
			ptr = ptr.Left
		}
	}

	return nil
}
