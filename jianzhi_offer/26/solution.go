package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if A.Val == B.Val && cmpTree(A, B) {
		return true
	}

	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func cmpTree(A, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	} else {
		return cmpTree(A.Left, B.Left) && cmpTree(A.Right, B.Right)
	}
}
