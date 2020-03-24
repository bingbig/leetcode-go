package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	nums := make([]int, k)
	getKNums(root, nums, k)
	return nums[0]
}

func getKNums(root *TreeNode, nums []int, k int) int {
	if root.Right != nil {
		k = getKNums(root.Right, nums, k)
	}

	if root != nil && k > 0 {
		k = k - 1
		nums[k] = root.Val
	}

	if root.Left != nil {
		k = getKNums(root.Left, nums, k)
	}
	return k
}

func addOn(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i]++
	}
}
